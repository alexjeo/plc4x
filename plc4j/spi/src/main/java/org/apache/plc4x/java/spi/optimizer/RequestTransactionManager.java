/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.spi.optimizer;

import java.util.Objects;
import java.util.Queue;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentLinkedQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Consumer;

/**
 * This is a limited Queue of Requests, a Protocol can use.
 * <p>
 * The Following Steps are necessary
 * <ul>
 *     <li>Register Slot</li>
 *     <li>Pass Runnable</li>
 *     <li>On Request or Response unregister Slot</li>
 * </ul>
 */
public class RequestTransactionManager {

    /** Executor that performs all operations */
    static final ExecutorService executor = Executors.newFixedThreadPool(4);
    private final Set<RequestTransaction> runningRequests;
    /** How many Transactions are allowed to run at the same time? */
    private int numberOfConcurrentRequests;
    /** Assigns each request a Unique Transaction Id, especially important for failure handling */
    private AtomicInteger transactionId = new AtomicInteger(0);
    /** Important, this is a FIFO Queue for Fairness! */
    private Queue<RequestTransaction> workLog = new ConcurrentLinkedQueue<>();

    public RequestTransactionManager(int numberOfConcurrentRequests) {
        this.numberOfConcurrentRequests = numberOfConcurrentRequests;
        // Immutable Map
        runningRequests = ConcurrentHashMap.newKeySet();
    }

    public RequestTransactionManager() {
        this(1);
    }

    public int getNumberOfConcurrentRequests() {
        return numberOfConcurrentRequests;
    }

    public void setNumberOfConcurrentRequests(int numberOfConcurrentRequests) {
        this.numberOfConcurrentRequests = numberOfConcurrentRequests;
    }

    public void submit(Consumer<RequestTransaction> context) {
        RequestTransaction transaction = startRequest();
        context.accept(transaction);
        // this.submit(transaction);
    }

    void submit(RequestTransaction handle) {
        assert handle.operation != null;
        // Add this Request with this handle i the Worklog
        // Put Transaction into Worklog
        this.workLog.add(handle);
        // Try to Process the Worklog
        processWorklog();
    }

    private void processWorklog() {
        while (runningRequests.size() < getNumberOfConcurrentRequests() && !workLog.isEmpty()) {
            RequestTransaction next = workLog.remove();
            this.runningRequests.add(next);
            Future<?> completionFuture = executor.submit(next.operation);
            next.setCompletionFuture(completionFuture);
        }
    }


    public RequestTransaction startRequest() {
        return new RequestTransaction(this, this.transactionId.getAndIncrement());
    }

    public int getNumberOfActiveRequests() {
        return this.runningRequests.size();
    }

    private void failRequest(RequestTransaction transaction) {
        // Try to fail it!
        transaction.getCompletionFuture().cancel(true);
        // End it
        endRequest(transaction);
    }

    private void endRequest(RequestTransaction transaction) {
        if (!this.runningRequests.contains(transaction)) {
            throw new IllegalArgumentException("Unknown Transaction or Transaction already finished!");
        }
        this.runningRequests.remove(transaction);
        // Process the worklog, a slot should be free now
        processWorklog();
    }

    public static class RequestTransaction {

        private final RequestTransactionManager parent;
        private final int transactionId;

        /** The iniital operation to perform to kick off the request */
        private Runnable operation;
        private Future<?> completionFuture;

        public RequestTransaction(RequestTransactionManager parent, int transactionId) {
            this.parent = parent;
            this.transactionId = transactionId;
        }

        public void start() {

        }

        public void failRequest(Throwable t) {
            this.parent.failRequest(this);
        }

        public void endRequest() {
            // Remove it from Running Requests
            this.parent.endRequest(this);
        }

        public void setOperation(Runnable operation) {
            this.operation = operation;
        }

        public Future<?> getCompletionFuture() {
            return completionFuture;
        }

        public void setCompletionFuture(Future<?> completionFuture) {
            this.completionFuture = completionFuture;
        }

        public void submit(Runnable operation) {
            this.setOperation(operation);
            this.parent.submit(this);
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            RequestTransaction that = (RequestTransaction) o;
            return transactionId == that.transactionId;
        }

        @Override
        public int hashCode() {
            return Objects.hash(transactionId);
        }
    }

}
