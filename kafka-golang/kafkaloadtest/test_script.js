import { check } from 'k6';
import { Writer, SchemaRegistry, SCHEMA_TYPE_JSON } from "k6/x/kafka";

const bootstrapServers = ['kafka-broker-1:9092'];
const kafkaTopic = 'notification';

function getRandomInt(max = 1000) {
    return Math.floor(Math.random() * max + 1);
}

export function produceMessage(message) {
    const writer = new Writer({
        brokers: bootstrapServers,
        topic: kafkaTopic,
    });

    writer.produce({ messages: [message] });

    writer.close();
}

export function createLoadKafkaMessage(accountId, orderKey) {
    const schemaRegistry = new SchemaRegistry();
    return {
        key: schemaRegistry.serialize({
            data: {
                key: accountId,
            },
            schemaType: SCHEMA_TYPE_JSON,
        }),
        value: schemaRegistry.serialize({ //TODO: This should move to a seperate directory, ideally we can send multiple different messages (not just copies of the same one)
            data: {
                "operation": "create",
                "order": {
                    "orderKey": `${orderKey}`,
                    "orderNumber": "123456",
                    "ownedByUserKey": "3a0060333f1d45ec807345967fa55c5",
                    "operatedByUserKey": "c323bce7df494cb490019f9dc8b21024",
                    "status": {
                        "code": "tendered",
                        "diagnostics": [
                            {
                                "message": "This is a diagnostic message - updated",
                                "propertyName": "currency",
                                "type": "insufficient_credit"
                            }
                        ],
                        "requestedChanges": [
                            {
                                "stagedChangeKey": "3a0060333f1d45ec807345967fa55ab8",
                                "type": "update",
                                "status": "pending",
                                "stagedChangeDate": "2023-11-29T18:37:29.684Z"
                            }
                        ]
                    },
                    "customerKey": "96a88a9ff20f4f0583d7431b356d5f26",
                    "money": {
                        "currency": "usd",
                        "tenderType": "spot",
                        "sellRateDetails": [
                            {
                                "rateCodeKey": "400",
                                "rate": 100,
                                "rateType": "distance",
                                "stopId": 1
                            },
                            {
                                "rateCodeKey": "405",
                                "rate": 100,
                                "rateType": "flat",
                                "multiplier": 1,
                                "stopId": 1
                            }
                        ],
                        "suggestedTopSpendAmount": 0
                    },
                    "stops": [
                        {
                            "stopId": 1,
                            "sequenceNumber": 1,
                            "customerLocationKey": null,
                            "type": "pick",
                            "name": "Test",
                            "addressLines": [
                                "600 W Chicago Ave."
                            ],
                            "city": "Chicago",
                            "stateOrProvince": "IL",
                            "postalCode": "60654",
                            "country": "us",
                            "blindAddress": null,
                            "contact": {
                                "customerLocationContactKey": null,
                                "name": "Jane Doe",
                                "role": "Warehouse coordinator",
                                "phoneNumberCountryCode": "1",
                                "phoneNumber": "5555555555",
                                "phoneExtension": "1234",
                                "faxCountryCode": "1",
                                "faxNumber": "1234567890",
                                "email": "bob@aws.com"
                            },
                            "serviceWindow": {
                                "startDate": "2023-11-27",
                                "startTime": "08:30",
                                "endDate": "2023-11-27",
                                "endTime": "15:30",
                                "timeZone": "America/Chicago"
                            },
                            "appointmentStatus": "required",
                            "locationType": "commercial",
                            "specialServices": null,
                            "internalNotes": "Internal test 1",
                            "externalNotes": "External test 1",
                            "billOfLadingNumber": "129",
                            "pickupDeliveryNumber": null,
                            "referenceValue": null
                        },
                        {
                            "stopId": 2,
                            "sequenceNumber": 2,
                            "customerLocationKey": null,
                            "type": "drop",
                            "name": "Test54",
                            "addressLines": [
                                "155 S SEWARD ST"
                            ],
                            "city": "JUNEAU",
                            "stateOrProvince": "AK",
                            "postalCode": "99801",
                            "country": "us",
                            "blindAddress": null,
                            "contact": {
                                "customerLocationContactKey": null,
                                "name": "Jane Doe",
                                "role": "Warehouse coordinator",
                                "phoneNumberCountryCode": "1",
                                "phoneNumber": "5555555555",
                                "phoneExtension": "1234",
                                "faxCountryCode": "1",
                                "faxNumber": "5555555555",
                                "email": "bob@aws.com"
                            },
                            "serviceWindow": {
                                "startDate": "2023-11-28",
                                "startTime": "08:30",
                                "endDate": "2023-11-28",
                                "endTime": "15:30",
                                "timeZone": "America/Chicago"
                            },
                            "appointmentStatus": "required",
                            "locationType": "commercial",
                            "specialServices": [
                                {
                                    "type": "liftgate_required",
                                    "quantity": null
                                }
                            ],
                            "internalNotes": "Internal test 2",
                            "externalNotes": "External test 2",
                            "billOfLadingNumber": null,
                            "pickupDeliveryNumber": null,
                            "referenceValue": null
                        }
                    ],
                    "distance": {
                        "value": 10,
                        "unit": "mile"
                    },
                    "equipment": [
                        "van_standard_48"
                    ],
                    "primaryContact": {
                        "name": "Jane Doe",
                        "phoneNumber": "5555555555",
                        "email": "bob@aws.com"
                    },
                    "driverSpecialServices": [
                        {
                            "type": "paps_fast"
                        },
                        {
                            "type": "blankets",
                            "quantity": 1
                        },
                        {
                            "type": "tarps",
                            "sizeInFeet": 4
                        }
                    ],
                    "sourcing": {
                        "mode": "truckload",
                        "shouldInitiateSourcing": true,
                        "isTargetedCommodity": true,
                        "isHighValue": true,
                        "directAssignUserKey": "12345",
                        "startAssignmentQueue": true,
                        "isAtRisk": true
                    },
                    "internalNotes": "Internal Notes",
                    "externalNotes": "External Notes",
                    "equipmentNotes": "Equipment Notes",
                    "createdByUserKey": "1e50dc3e5aea43018bd6960fac76cabb",
                    "submittedByUserKey": "1e50dc3e5aea43018bd6960fac76cabb",
                    "createdBy": {
                        "userKey": "1e50dc3e5aea43018bd6960fac76cabb",
                        "systemAccountKey": "1e50dc3e5aea43018bd6960fac76cabb"
                    },
                    "submittedBy": {
                        "userKey": "1e50dc3e5aea43018bd6960fac76cabb",
                        "systemAccountKey": "1e50dc3e5aea43018bd6960fac76cabb"
                    },
                    "lastModifiedBy": {
                        "userKey": "1e50dc3e5aea43018bd6960fac76cabb",
                        "systemAccountKey": "1e50dc3e5aea43018bd6960fac76cabb"
                    },
                    "createdDate": "2023-11-29T18:37:44.592Z",
                    "lastModifiedDate": "2023-11-29T18:37:44.592Z",
                }
            }
            ,
            schemaType: SCHEMA_TYPE_JSON,
        }),
    }
}

export default function () {

    const message = createLoadKafkaMessage("abc", getRandomInt);
    produceMessage(message);
}