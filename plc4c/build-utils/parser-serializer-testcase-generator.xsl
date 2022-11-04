<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:xs="http://www.w3.org/2001/XMLSchema"
                xmlns:test="https://plc4x.apache.org/schemas/parser-serializer-testsuite.xsd"
                exclude-result-prefixes="xs"
                version="2.0">

    <xsl:output omit-xml-declaration="yes" />

    <xsl:template match="/test:testsuite">
        <xsl:variable name="protocolName" select="protocolName/text()"/>
        <xsl:variable name="outputFlavor" select="outputFlavor/text()"/>
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

// Code generated by code-generation. DO NOT EDIT.

#include <xsl:value-of disable-output-escaping="yes" select="'&#x3C;'"/>unity.h<xsl:value-of disable-output-escaping="yes" select="'&#x3E;'"/>

#include "plc4c/spi/read_buffer.h"
#include "plc4c/spi/write_buffer.h"
<xsl:for-each-group select="/test:testsuite/testcase/root-type" group-by="text()">
    <xsl:variable name="rootTypeName">
        <xsl:call-template name="getRootTypeName">
            <xsl:with-param name="name" select="current-grouping-key()"/>
        </xsl:call-template>
    </xsl:variable>
#include "<xsl:value-of select="$rootTypeName"/>.h"
</xsl:for-each-group>

void internal_assert_arrays_equal(uint8_t* expected_array, uint8_t* actual_array, uint8_t num_bytes);

        <xsl:for-each select="testcase">
            <xsl:variable name="testName">
                <xsl:call-template name="getTestName">
                    <xsl:with-param name="protocolName" select="$protocolName"/>
                    <xsl:with-param name="outputFlavor" select="$outputFlavor"/>
                    <xsl:with-param name="testNode" select="."/>
                </xsl:call-template>
            </xsl:variable>
            <xsl:variable name="rootTypeName">
                <xsl:call-template name="getRootTypeName">
                    <xsl:with-param name="name" select="normalize-space(root-type)"/>
                </xsl:call-template>
            </xsl:variable>
void <xsl:value-of select="normalize-space($testName)"/>() {
            <xsl:variable name="testData">
                <xsl:call-template name="getTestData">
                    <xsl:with-param name="rawData" select="raw"/>
                </xsl:call-template>
            </xsl:variable>
    uint8_t payload[] = {
        <xsl:value-of select="$testData"/>
    };
    uint16_t payload_size = sizeof(payload);

    // Create a new read_buffer instance
    plc4c_spi_read_buffer* read_buffer;
    plc4c_return_code return_code =
    plc4c_spi_read_buffer_create(payload, payload_size, <xsl:value-of disable-output-escaping="yes" select="'&#038;'"/>read_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error creating read buffer");
    }

    plc4c_<xsl:value-of select="replace(lower-case($protocolName), ' ', '_')"/>_<xsl:value-of select="replace(replace(lower-case($outputFlavor), ' ', '_'), '-', '_')"/>_<xsl:value-of select="$rootTypeName"/>* message = NULL;
    return_code = plc4c_<xsl:value-of select="replace(lower-case($protocolName), ' ', '_')"/>_<xsl:value-of select="replace(replace(lower-case($outputFlavor), ' ', '_'), '-', '_')"/>_<xsl:value-of select="$rootTypeName"/>_parse(read_buffer, <xsl:value-of disable-output-escaping="yes" select="'&#038;'"/>message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error parsing packet");
    }

    plc4c_spi_write_buffer* write_buffer;
    return_code = plc4c_spi_write_buffer_create(payload_size, <xsl:value-of disable-output-escaping="yes" select="'&#038;'"/>write_buffer);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error writing to buffer");
    }

    return_code = plc4c_<xsl:value-of select="replace(lower-case($protocolName), ' ', '_')"/>_<xsl:value-of select="replace(replace(lower-case($outputFlavor), ' ', '_'), '-', '_')"/>_<xsl:value-of select="$rootTypeName"/>_serialize(write_buffer, message);
    if (return_code != OK) {
        TEST_FAIL_MESSAGE("Error serializing");
    }

    internal_assert_arrays_equal(payload, write_buffer-<xsl:value-of disable-output-escaping="yes" select="'&#x3E;'"/>data, payload_size);

    printf("Success");
}
        </xsl:for-each>


void parser_serializer_test_<xsl:value-of select="replace(lower-case($protocolName), ' ', '_')"/>_<xsl:value-of select="replace(replace(lower-case($outputFlavor), ' ', '_'), '-', '_')"/>(void) {
        <xsl:for-each select="testcase">
            <xsl:variable name="testName">
                <xsl:call-template name="getTestName">
                    <xsl:with-param name="protocolName" select="$protocolName"/>
                    <xsl:with-param name="outputFlavor" select="$outputFlavor"/>
                    <xsl:with-param name="testNode" select="."/>
                </xsl:call-template>
            </xsl:variable>
    RUN_TEST(<xsl:value-of select="$testName"/>);
        </xsl:for-each>
}
    </xsl:template>
    
    <xsl:template name="getTestName">
        <xsl:param name="protocolName"/>
        <xsl:param name="outputFlavor"/>
        <xsl:param name="testNode"/>
        parser_serializer_test_<xsl:value-of select="replace(lower-case($protocolName), ' ', '_')"/>_<xsl:value-of select="replace(replace(lower-case($outputFlavor), ' ', '_'), '-', '_')"/>_<xsl:value-of select="translate(replace(lower-case($testNode/name), ' ', '_'), '()', '')"/>
    </xsl:template>
    
    <xsl:template name="getRootTypeName">
        <xsl:param name="name"/>
        <xsl:analyze-string select="$name" regex="(.+)([A-Z][a-z].*)">
            <xsl:matching-substring>
                <xsl:variable name="restName">
                    <xsl:call-template name="getRootTypeName">
                        <xsl:with-param name="name" select="regex-group(2)"/>
                    </xsl:call-template>
                </xsl:variable>
                <xsl:value-of select="lower-case(regex-group(1))"/>_<xsl:value-of select="$restName"/>
            </xsl:matching-substring>
            <xsl:non-matching-substring>
                <xsl:value-of select="lower-case($name)"/>
            </xsl:non-matching-substring>
        </xsl:analyze-string>
    </xsl:template>

    <xsl:template name="getTestData">
        <xsl:param name="rawData"/>
        <xsl:variable name="normalizedRawData" select="normalize-space($rawData)"/>
        <xsl:for-each select="0 to (string-length($normalizedRawData) - 1) idiv 2">0x<xsl:value-of select="substring($normalizedRawData, (. * 2) + 1, 2)"/><xsl:if test="position() != last()">, </xsl:if></xsl:for-each>
    </xsl:template>

</xsl:stylesheet>