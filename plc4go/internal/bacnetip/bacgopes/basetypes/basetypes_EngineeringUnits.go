/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package basetypes

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type EngineeringUnits struct {
	*Enumerated
	vendorRange  vendorRange
	enumerations map[string]uint64
}

func NewEngineeringUnits(arg Arg) (*EngineeringUnits, error) {
	s := &EngineeringUnits{
		vendorRange: vendorRange{256, 65535},
		enumerations: map[string]uint64{
			//Acceleration
			"metersPerSecondPerSecond": 166,
			"squareMeters":             0,
			"squareCentimeters":        116,
			"squareFeet":               1,
			"squareInches":             115,
			//Currency
			"currency1":  105,
			"currency2":  106,
			"currency3":  107,
			"currency4":  108,
			"currency5":  109,
			"currency6":  110,
			"currency7":  111,
			"currency8":  112,
			"currency9":  113,
			"currency10": 114,
			//Electrical
			"milliamperes":                2,
			"amperes":                     3,
			"amperesPerMeter":             167,
			"amperesPerSquareMeter":       168,
			"ampereSquareMeters":          169,
			"decibels":                    199,
			"decibelsMillivolt":           200,
			"decibelsVolt":                201,
			"farads":                      170,
			"henrys":                      171,
			"ohms":                        4,
			"ohmMeters":                   172,
			"ohmMeterPerSquareMeter":      237,
			"milliohms":                   145,
			"kilohms":                     122,
			"megohms":                     123,
			"microSiemens":                190,
			"millisiemens":                202,
			"siemens":                     173,
			"siemensPerMeter":             174,
			"teslas":                      175,
			"volts":                       5,
			"millivolts":                  124,
			"kilovolts":                   6,
			"megavolts":                   7,
			"voltAmperes":                 8,
			"kilovoltAmperes":             9,
			"megavoltAmperes":             10,
			"ampereSeconds":               238,
			"ampereSquareHours":           246,
			"voltAmpereHours":             239, //VAh
			"kilovoltAmpereHours":         240, //kVAh
			"megavoltAmpereHours":         241, //MVAh
			"voltAmperesReactive":         11,
			"kilovoltAmperesReactive":     12,
			"megavoltAmperesReactive":     13,
			"voltAmpereHoursReactive":     242, //varh
			"kilovoltAmpereHoursReactive": 243, //kvarh
			"megavoltAmpereHoursReactive": 244, //Mvarh
			"voltsPerDegreeKelvin":        176,
			"voltsPerMeter":               177,
			"voltsSquareHours":            245,
			"degreesPhase":                14,
			"powerFactor":                 15,
			"webers":                      178,
			// Energy
			"joules":                16,
			"kilojoules":            17,
			"kilojoulesPerKilogram": 125,
			"megajoules":            126,
			"joulesPerHours":        247,
			"wattHours":             18,
			"kilowattHours":         19,
			"megawattHours":         146,
			"wattHoursReactive":     203,
			"kilowattHoursReactive": 204,
			"megawattHoursReactive": 205,
			"btus":                  20,
			"kiloBtus":              147,
			"megaBtus":              148,
			"therms":                21,
			"tonHours":              22,
			// Enthalpy
			"joulesPerKilogramDryAir":     23,
			"kilojoulesPerKilogramDryAir": 149,
			"megajoulesPerKilogramDryAir": 150,
			"btusPerPoundDryAir":          24,
			"btusPerPound":                117,
			"joulesPerDegreeKelvin":       127,
			// Entropy
			"kilojoulesPerDegreeKelvin":     151,
			"megajoulesPerDegreeKelvin":     152,
			"joulesPerKilogramDegreeKelvin": 128,
			// Force
			"newton": 153,
			// Frequency
			"cyclesPerHour":                   25,
			"cyclesPerMinute":                 26,
			"hertz":                           27,
			"kilohertz":                       129,
			"megahertz":                       130,
			"perHour":                         131,
			"gramsOfWaterPerKilogramDryAir":   28,
			"percentRelativeHumidity":         29,
			"micrometers":                     194,
			"millimeters":                     30,
			"centimeters":                     118,
			"kilometers":                      193,
			"meters":                          31,
			"inches":                          32,
			"feet":                            33,
			"candelas":                        179,
			"candelasPerSquareMeter":          180,
			"wattsPerSquareFoot":              34,
			"wattsPerSquareMeter":             35,
			"lumens":                          36,
			"luxes":                           37,
			"footCandles":                     38,
			"milligrams":                      196,
			"grams":                           195,
			"kilograms":                       39,
			"poundsMass":                      40,
			"tons":                            41,
			"gramsPerSecond":                  154,
			"gramsPerMinute":                  155,
			"kilogramsPerSecond":              42,
			"kilogramsPerMinute":              43,
			"kilogramsPerHour":                44,
			"poundsMassPerSecond":             119,
			"poundsMassPerMinute":             45,
			"poundsMassPerHour":               46,
			"tonsPerHour":                     156,
			"milliwatts":                      132,
			"watts":                           47,
			"kilowatts":                       48,
			"megawatts":                       49,
			"btusPerHour":                     50,
			"kiloBtusPerHour":                 157,
			"horsepower":                      51,
			"tonsRefrigeration":               52,
			"pascals":                         53,
			"hectopascals":                    133,
			"kilopascals":                     54,
			"pascalSeconds":                   253,
			"millibars":                       134,
			"bars":                            55,
			"poundsForcePerSquareInch":        56,
			"millimetersOfWater":              206,
			"centimetersOfWater":              57,
			"inchesOfWater":                   58,
			"millimetersOfMercury":            59,
			"centimetersOfMercury":            60,
			"inchesOfMercury":                 61,
			"degreesCelsius":                  62,
			"degreesKelvin":                   63,
			"degreesKelvinPerHour":            181,
			"degreesKelvinPerMinute":          182,
			"degreesFahrenheit":               64,
			"degreeDaysCelsius":               65,
			"degreeDaysFahrenheit":            66,
			"deltaDegreesFahrenheit":          120,
			"deltaDegreesKelvin":              121,
			"years":                           67,
			"months":                          68,
			"weeks":                           69,
			"days":                            70,
			"hours":                           71,
			"minutes":                         72,
			"seconds":                         73,
			"hundredthsSeconds":               158,
			"milliseconds":                    159,
			"newtonMeters":                    160,
			"millimetersPerSecond":            161,
			"millimetersPerMinute":            162,
			"metersPerSecond":                 74,
			"metersPerMinute":                 163,
			"metersPerHour":                   164,
			"kilometersPerHour":               75,
			"feetPerSecond":                   76,
			"feetPerMinute":                   77,
			"milesPerHour":                    78,
			"cubicFeet":                       79,
			"cubicFeetPerDay":                 248,
			"cubicMeters":                     80,
			"cubicMetersPerDay":               249,
			"imperialGallons":                 81,
			"milliliters":                     197,
			"liters":                          82,
			"usGallons":                       83,
			"cubicFeetPerSecond":              142,
			"cubicFeetPerMinute":              84,
			"cubicFeetPerHour":                191,
			"cubicMetersPerSecond":            85,
			"cubicMetersPerMinute":            165,
			"cubicMetersPerHour":              135,
			"imperialGallonsPerMinute":        86,
			"millilitersPerSecond":            198,
			"litersPerSecond":                 87,
			"litersPerMinute":                 88,
			"litersPerHour":                   136,
			"usGallonsPerMinute":              89,
			"usGallonsPerHour":                192,
			"degreesAngular":                  90,
			"degreesCelsiusPerHour":           91,
			"degreesCelsiusPerMinute":         92,
			"degreesFahrenheitPerHour":        93,
			"degreesFahrenheitPerMinute":      94,
			"jouleSeconds":                    183,
			"kilogramsPerCubicMeter":          186,
			"kilowattHoursPerSquareMeter":     137,
			"kilowattHoursPerSquareFoot":      138,
			"megajoulesPerSquareMeter":        139,
			"megajoulesPerSquareFoot":         140,
			"noUnits":                         95,
			"newtonSeconds":                   187,
			"newtonsPerMeter":                 188,
			"partsPerMillion":                 96,
			"partsPerBillion":                 97,
			"percent":                         98,
			"percentObscurationPerFoot":       143,
			"percentObscurationPerMeter":      144,
			"percentPerSecond":                99,
			"perMinute":                       100,
			"perSecond":                       101,
			"psiPerDegreeFahrenheit":          102,
			"radians":                         103,
			"radiansPerSecond":                184,
			"revolutionsPerMinute":            104,
			"squareMetersPerNewton":           185,
			"wattsPerMeterPerDegreeKelvin":    189,
			"wattsPerSquareMeterDegreeKelvin": 141,
			"perMille":                        207,
			"gramsPerGram":                    208,
			"kilogramsPerKilogram":            209,
			"gramsPerKilogram":                210,
			"milligramsPerGram":               211,
			"milligramsPerKilogram":           212,
			"gramsPerMilliliter":              213,
			"gramsPerLiter":                   214,
			"milligramsPerLiter":              215,
			"microgramsPerLiter":              216,
			"gramsPerCubicMeter":              217,
			"milligramsPerCubicMeter":         218,
			"microgramsPerCubicMeter":         219,
			"nanogramsPerCubicMeter":          220,
			"gramsPerCubicCentimeter":         221,
			"wattHoursPerCubicMeter":          250,
			"joulesPerCubicMeter":             251,
			"becquerels":                      222,
			"kilobecquerels":                  223,
			"megabecquerels":                  224,
			"gray":                            225,
			"milligray":                       226,
			"microgray":                       227,
			"sieverts":                        228,
			"millisieverts":                   229,
			"microsieverts":                   230,
			"microsievertsPerHour":            231,
			"decibelsA":                       232,
			"nephelometricTurbidityUnit":      233,
			"pH":                              234,
			"gramsPerSquareMeter":             235,
			"minutesPerDegreeKelvin":          236,
		},
	}
	panic("enumeratedimplementme")
	return s, nil
}