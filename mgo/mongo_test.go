package mgo

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/rules"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	//update := bson.D{{"$set",bson.A{}}}
	//log.Print(len(update))
	Connect("localhost:27017", "ricnsmart_dev")

	filter := bson.M{"deviceid": "6d8a69ce-8b4d-4caf-be55-eaa2c828bfaa"}
	//filter := bsonx.Doc{{"DeviceID", bsonx.String("6c76660a-d285-48b7-b4c8-ac7a70179153")}}
	//update := bsonx.Doc{{"$set", bsonx.Document(bsonx.Doc{{"Metric.Ia.Alert", bsonx.Int32(80)}})}}

	update := bson.D{{"$set", bson.M{"metrics.ia": bson.M{
		"smsswitch":   false,
		"warnswitch":  false,
		"alertswitch": false,
		"max":         300,
		"min":         0,
		"scale":       5,
		"warn":        62,
		"alert":       82,
		"spl":         12,
	}}}}

	_, err := MongoDB.Collection(rules.DevicesCollection).UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

}

func TestDecode(t *testing.T) {
	var deviceMetric DeviceMetric

	Connect("localhost:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"deviceid": "6d8a69ce-8b4d-4caf-be55-eaa2c828bfaa"}

	err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).Decode(&deviceMetric)
	//b, err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(`%+v`, deviceMetric)
	//log.Print(b.String())

}

func TestOR(t *testing.T) {
	Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	a := bson.A{}

	a = append(a, bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}, bson.M{"deviceid": "aabb6480-6412-485a-9397-7277f110704d"})

	filter := bson.D{
		{"$or", a},
	}

	//filter := bson.M{"deviceid": "1354b470-7733-4194-b406-f9a5999d8d57"}

	//b, err := MongoDB.Collection(rules.DevicesCollection).FindOne(ctx, filter).DecodeBytes()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Print(b.String())

	//
	cur, err := MongoDB.Collection(rules.DevicesCollection).Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var deviceMetrics []DeviceMetric
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var deviceMetric DeviceMetric
		err := cur.Decode(&deviceMetric)
		if err != nil {
			log.Fatal(err)
		}
		deviceMetrics = append(deviceMetrics, deviceMetric)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", deviceMetrics)

}

func TestDelete(t *testing.T) {
	Connect("mongodb://39.104.186.37:27017", "ricnsmart_dev")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	arr := bson.A{}

	for _, deviceID := range []string{"1354b470-7733-4194-b406-f9a5999d8d57", "aabb6480-6412-485a-9397-7277f110704d"} {
		arr = append(arr, bson.M{"deviceid": deviceID})
	}

	filter := bson.D{{"$or", arr}}

	_, err := MongoDB.Collection(rules.DevicesCollection).DeleteMany(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
}

type DeviceMetric struct {
	GPRSOperator int
	DomainRecord string
	CT           int
	Interval     int
	SMSLimit     int
	Metrics      struct {
		DO1 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DO2 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DI1 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		DI2 struct {
			AlertSwitch bool
			SMSSwitch   bool
		}
		Ia struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
			SPL         int
		}
		Ib struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
			SPL         int
		}
		Ic struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
			SPL         int
		}
		IR struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Warn        int
			Alert       int
			SPL         int
		}
		T1 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
		T2 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
		T3 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
		T4 struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Warn        int
			Alert       int
		}
		Ua struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
		Ub struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
		Uc struct {
			WarnSwitch  bool
			AlertSwitch bool
			SMSSwitch   bool
			Min         int
			Max         int
			Scale       int
			Warn        int
			Alert       int
		}
	}
}

func TestMigration(t *testing.T) {
	Connect("mongodb://39.104.158.136:27017", "ricnsmart")

	var docs []interface{}

	//var origin = bson.M{}

	arr := []string{
		"3697a8cd-6fc2-4416-ad0a-f45434756b75",
		"5b11af6e-3514-463d-82cc-d80d92f91979",
		"19580516-3cd2-4389-849a-53e6dcc82bc5",
		"83da5d95-3fb3-450f-b1e8-72da66a977d7",
		"14d4ef3c-f607-462b-983d-0f63ead77135",
		"b90fdd72-79b1-4bf2-ac75-156f487902eb",
		"331e2528-9ea7-4422-9996-faea4a92f0f3",
		"4933aa23-e886-45d2-b1ec-121dad683dcb",
		"a79eb19c-5655-486f-a0a8-844596257309",
		"4353b065-a7a3-43b7-a61c-c2e6681a3132",
		"2ec4e98a-326c-414a-a65f-c04a8aad3c5b",
		"1a0017ca-ce8d-4f74-a9da-02dc7799b344",
		"af7e597e-c8ac-4862-983b-817390df2426",
		"695c6730-37bf-4a16-8c55-b4191f98dcaa",
		"0b6419d7-d35e-49be-8c8b-6b981edc765e",
		"8690cc82-9a00-42e3-9824-beb7aed851ee",
		"69576c52-78fb-4d41-b88e-dda1eba3ffe5",
		"fe74809f-86a0-44fb-8290-dd72e1e25a2d",
		"4d2b5f5e-fe50-41b3-b6fe-d3e1e11e91f1",
		"8ab560b2-4a4c-4593-9c06-5e4590c3546d",
		"bdce7bc8-27a7-4d6c-8a18-e1dd3472964e",
		"c40c1bda-baa5-40e3-9cf6-ae3c02a350a0",
		"6d90fd33-b9fc-49fc-86ad-4d6db0a21068",
		"0293c3f5-8bdc-4c45-923a-2f7e03087afd",
		"34ba8a69-0626-47c8-9186-10e5b326da23",
		"ff0c668d-b1df-41bb-a896-07676103964c",
		"0661d0fa-1f80-4de8-989b-4c858533b2b3",
		"7bb2ceff-404b-4ddd-9089-3605098bad30",
		"e722a5f2-e547-4b0d-9106-3561a8e5bfad",
		"81859ce3-8a6a-4227-a9e9-5eaa4a5af55f",
		"a13734f4-0256-40e7-aa48-72e9b156efbb",
		"6d8a69ce-8b4d-4caf-be55-eaa2c828bfaa",
		"5253f818-7ea5-4a25-ab3a-3a9f8828b1f6",
		"9aa17a74-005a-4c7f-a564-0c3e0b9e5be9",
		"48d59192-9696-4b09-af52-5a1ff7f58946",
		"8ecf2a21-e459-49a2-a98a-613e68c1db76",
		"bc30af1f-e0bc-49c6-b716-3175e82962ea",
		"8609550b-78c0-44dc-99c0-8052efac597a",
		"7ced0c3b-2992-4d89-b81f-7e8824aaa716",
		"93973e71-d38a-4951-92ad-a96421fda840",
		"297ed247-8e3b-4f8a-9bd8-07d1104164cb",
		"9cad589a-51d0-4fe1-a8eb-ec17090e7980",
		"984e8a5f-7db2-4ac9-8972-f72536c562f5",
		"d1f93831-c03e-4270-acb7-78c01926d7b7",
		"815a3cbf-fd2b-4c78-8e02-84a1938fdb73",
		"74c799c3-bafa-4717-88f4-d6f44a738021",
		"4112446a-0278-43e5-a717-7c5f25458837",
		"6bd47321-0bae-4e33-99f5-be790f60cbcb",
		"7b2f5a3e-03cf-404d-8a4e-c7aaa5e32680",
		"e1c25fb4-62bf-435c-b690-808ffb037c25",
		"d3521bec-dd0d-4372-8297-785c062b481b",
		"bb8f539c-ecf2-4441-b435-1091981e7fa8",
		"d02f370d-69ea-40d3-83f1-33c2487bf1ae",
		"cf8b7685-4df3-4006-bfa7-b793590f59ac",
		"6b9ce8aa-7fa3-453c-b682-6b21344f8f1a",
		"cfee6999-98bb-43e5-b570-70331dd50496",
		"d12deb02-eb55-4aee-bc7d-b6525d9683d0",
		"c27e6a52-d902-4381-bd98-1dc2e195def7",
		"ddb44430-d8f3-4587-8f37-e3b0ce5807db",
		"deeee4c3-3655-4ba0-a9b6-c9f5f7dcc232",
		"d80ec50c-6032-4980-bd2a-0e0a6bfbcb69",
		"8429821b-fb2f-4417-ae83-e033a74a57ba",
		"3fc7bc6d-2e4d-4e86-843f-f1882e545698",
		"ccf2df31-89bb-45d9-9f1d-97d9d8183dcb",
		"64bb9f8b-d7d1-4a26-ba60-34e026fe143c",
		"8bb6e049-7750-4f43-b308-0218fd8dd399",
		"368345c2-a39e-447a-88a8-eb25285c5616",
		"5e2cb3d6-cd1d-415a-934e-8f4974a2633f",
		"0ecee8fe-d82f-43db-8f21-d6b9fd906fb7",
		"850cb293-47be-4fcb-973b-3359715ee3bd",
		"74625006-7d41-4476-bd56-de33f6f79159",
		"9f3b306e-f5f0-4f5e-95bb-9f0fe5d6b537",
		"2b4cb033-a409-4844-8315-b5e45e6ea698",
		"4c01d040-2edd-4de1-89df-0afa5e1562ff",
		"21e40236-2ba8-4523-96ec-1f9b1c1db8e0",
		"f4b959ec-ae21-4d19-9734-c0a4eb8b1e86",
		"74b69c4e-8ca9-4326-a35a-b6cd443554a6",
		"0d4a80c3-b845-4452-b518-c013ff9c6277",
		"2596a5f2-a5d0-422d-9ebb-f818f7321fdc",
		"736fc6f8-fb2d-421c-95e6-8437342700dd",
		"f83ab011-ab00-49c6-9cb1-132210cea961",
		"ddac1b2c-c15d-415a-8ef3-add4edd53f9e",
		"4ddb7b63-7166-4b94-9a9a-599dd0ee3732",
		"8462d151-acdf-4e30-a357-48eaf4dff69f",
		"392e022e-1b67-4dc0-abe3-fdaade6007a2",
		"bb09e34d-12a2-4045-9687-de219bf1470f",
		"8efb381a-5417-4827-9842-38ce6e736f88",
		"202e5040-dafc-4815-bd00-ed5e00a937a1",
		"5922bac9-8afa-41b4-8921-ceff7086851b",
		"25ebb540-fd23-4e3e-95b2-5aa2b1b03ecf",
		"9c8a3ec1-32b6-46f8-a7f3-59441c560f57",
		"be84ff4d-d8bd-4477-9c0b-874bfee4f934",
		"03b2be0a-f72f-4822-9228-80d2076688e5",
		"9be034ca-1031-4677-8a10-a9b175dd4e11",
		"c567710e-f444-4b3a-9a50-56ae4c2122bd",
		"206276f4-aca1-4f28-9759-8f503fa8411e",
		"b329da02-4fb4-40f1-abd6-746570cdc0bf",
		"50263418-0877-4cbe-b5cb-c1687dea7361",
		"33d94157-a0cf-45cb-adaa-d5643261f1fe",
		"6763d419-f137-43af-9f5e-418e17a9a7bf",
		"48b9da98-a2e1-4f55-b14e-51a0a054ad50",
		"5dfaa86c-93a8-4223-bc8a-9b121d39f880",
		"a261e12f-09e4-4ee0-94da-0929f3769017",
		"35946906-4df0-4853-9d80-ec57a82954f6",
		"f69d75b2-61b8-4d67-b3b5-5b37c27cb24e",
		"baef1545-f639-41f2-bdb8-a15025b5fbde",
		"605707d7-646b-4bca-abe9-78f2a7170068",
		"3399bab4-14f7-496f-b395-5a10aeb62dfb",
		"6d34dc83-fc66-4d62-82af-d572c5fdb531",
		"5dd3e0f7-92c0-469f-8b2e-519b5605fc9d",
		"d407d1fb-4fdd-4341-9540-871e0cda8697",
		"c1ae0123-0c23-4ab1-85dd-af68d72cd9f9",
		"10ca52ff-1b06-4825-ad0c-3cc3ca65f3ed",
		"7ac23a8f-c525-4fe7-87e6-a7de959cb4ab",
		"94b0b400-c7b1-4761-beab-35ebb861dd31",
		"6b2a627f-00d4-4c4b-97fc-08ac2ac68a67",
		"185acdec-52e0-4a76-bc9a-6be1b9378c80",
		"92619ae1-53b8-4415-af99-fef34c597d46",
		"2535e0d7-ec59-4206-a246-b7f1eb2432b4",
		"26714cac-fed0-429c-908a-15fc4a69a516",
		"d650c12f-0981-4cce-979b-641006fa7437",
		"068e35b2-cfc9-4720-b857-a8c8f858f4df",
		"7b5c2683-eb3d-4ff7-9a30-ee133fef08fd",
		"a4e4582f-4a00-4916-9e28-91a97134a8c5",
		"a081080b-c3a8-48f4-9cff-63ce309d0619",
		"bbef10ec-6589-46a6-8a98-6785f80bc9c1",
		"af3e3d4f-3749-4a30-b4b2-c6a5ff4db015",
		"0807a507-5535-47e9-998d-ec6633d3226a",
		"b5137d16-fd9b-4135-8cb0-f6d0badee66a",
		"54ffd8cf-0032-415a-ae88-30ebf0560edc",
		"2f1bf8ea-2ab3-49fa-877d-e364c058121d",
		"39f62327-ee2c-450f-a7f0-c8edb0ea2379",
		"9c26477e-55b1-4449-8928-5ed4788abe8c",
		"902261dd-56fa-458b-a996-b9db6423df4b",
		"615a5090-f546-4150-b1c4-5da2c1ca269b",
		"6088e595-336e-4890-aad8-537c8069e460",
		"a074e35d-90d7-4cf7-b3e3-aeb9b2fb1ac2",
		"536a98ce-34c0-42c7-ae9f-4795e5329392",
		"b7f9660c-dfb1-4872-ba5b-85730692627a",
		"b69653c8-c5d7-4848-84c0-947b7d923e21",
		"6b4a8f8f-d5d3-40e7-9d06-4af1446fc238",
		"ac009b67-0228-48dc-9477-183b0d4526cb",
		"f5b7700d-964f-4530-a715-18a85ed43fab",
		"5e43b01d-0005-43f5-8bcc-340b666e9104",
		"56518a41-7736-41a3-bdce-da5e2a999a3b",
		"d89f77ff-433c-4d0e-9041-6366825d321d",
		"de11e8cb-5139-4b0f-9efb-d0632da75341",
		"ba260a8c-734d-4d9a-a3fc-bba79a57a35b",
		"7ff7deb0-51bc-4621-890c-2e1e93d5e82e",
		"35564c07-9daa-4e00-b45c-c01288932e0d",
		"fca24f5a-345f-4e80-9e69-7de7a2bf0f2a",
		"eb58450a-0d17-4e71-af2b-70d50f09f356",
		"79708e73-d6ba-4f3d-bad8-e263f58fcfc9",
		"10ce668d-2e87-48cb-86d0-9cfdda0d5915",
		"b0095664-3b57-41b5-961d-072824ab6ed5",
		"75662cdd-edad-4740-a8b3-07187e4d153f",
		"89438731-a20f-4551-8201-b04fd5caf136",
		"b5160a29-c640-40e3-9510-ba1b655a6f3e",
		"6daad684-1aad-4b68-8b52-a3cd3c74a6e3",
		"e30ee36d-6443-467f-be48-01d2033b9cd9",
		"393f0429-ed30-4c4b-b7d4-56de590a859e",
		"65aa17b8-8ddb-4cc9-b6cd-f5194f0028be",
		"cbf3d733-5cbb-4fa1-afe3-3b9bfbd7cb28",
		"22a4ff45-1344-462b-850e-a0d797173cef",
		"df070184-23ab-41fc-b2b4-d76a828b2b5d",
		"c0c2195e-07d5-4a46-8038-21078a0523bf",
		"01ac47a9-40ec-4bb1-ba29-a7238fe47017",
		"0c27028e-7b0e-450a-a4ab-64ad1329f033",
		"b8a8322c-1cbe-437f-bd65-8a3fa0ffb364",
		"aa58ed14-6f1b-482b-9d2c-690ed73a80e4",
		"7cc8dddc-7f99-411c-9f63-0655de9d2650",
		"ed3374a4-9838-4241-ba68-897819fd5749",
		"872bb438-11ab-4ffe-ae4a-03b9575a1be8",
		"f728a0e8-7c2b-4700-8cff-0a858df7b4e6",
		"646c3bc2-2367-4447-8cb7-dd84ad6c801f",
		"45a09349-306c-4b6a-8c90-da1905291a36",
		"3021e2ff-b746-49e2-9dfe-2e4d72461ac8",
		"cff3eba3-9b26-4a76-9d4c-59250eb55fec",
		"999d1b90-7908-4767-acfd-1aa2c45dcd5a",
		"b85594bc-a615-40a4-b675-4b2ef99fec4f",
		"5b0ecd4a-04d2-4aae-b553-12ccef0ba7eb",
		"38e24a00-9451-43f6-9839-6778247a29ea",
		"edfc4aa3-fa02-4310-a960-271781ed532b",
		"12cad7bf-705e-47cd-b6c1-e21f259c111e",
		"f7704f71-44d9-4ef3-9ffb-f900b51a5a19",
		"d7c992da-012e-41e1-83f2-999b4e99aec2",
		"0f99a503-454c-49fc-9404-d77cffd7a23a",
		"c4e0e96e-0ee7-4242-8a21-32478a46329e",
		"2559b47e-3d4f-41bb-9d82-25aa026e68fe",
		"66a2272a-96d9-45a9-8167-0d14c0168542",
		"ff8a57b6-1edb-4b1e-a58b-84ecfb7551fd",
		"a4e85f94-7705-407f-b4d0-510857de399b",
		"2f778e0f-2b07-4937-aaf7-aecccd9940e1",
		"0e99522f-f295-4f6e-850a-b760256047bd",
		"a5293d81-4137-4ca9-809d-5433b5463ce8",
		"6f66a36b-adc5-467c-a955-4e19d26dda73",
		"357aed50-eb22-4119-8343-1503457f3bf8",
		"9441a465-5f25-4afb-9840-095259eb0ad9",
		"be4caab3-78a7-4788-b066-8f35c111f351",
		"4091a268-c3c7-49a2-a8dc-22d9412a8eb0",
		"43e47477-8781-4da9-acfa-ddebec8bf682",
		"3c19af3d-65fb-4ee7-b3c5-42895052329a",
		"3a3b9f81-1f8d-49f7-9384-e13f65dba003",
		"80024c4a-76f8-40a4-87cf-9b96f9ff73d4",
		"49676f25-ae1d-44ca-af51-8b812726ceb6",
		"f060d51b-4df1-4dcf-9678-618bc6d266d5",
		"77b5c4cc-f6ca-4626-814f-ad7331575f14",
		"8ee99db6-74b1-46df-9c15-13cd54387a06",
		"cd387499-a511-46b3-aadd-a44018b119c2",
		"b1805762-b90b-4a98-88fd-e98959284364",
		"4691f581-727d-434d-bc1c-10a6be3ebe8f",
		"d5b1fcf0-2093-4340-8e08-5152e06ccac7",
		"9b2358de-eaeb-4ee2-8f7f-959a5f80d829",
		"1db02a48-d8d5-4b89-83fa-ae7aeba1a3bb",
		"57da5d47-e948-417c-9211-8de82078b1d3",
		"2b2f5aa0-1c95-4b61-9381-104fc4df884d",
		"867b030a-273a-4bae-8760-2fea0cc39599",
		"d9ef73e0-43d6-47b8-82c8-0825bef8c2be",
		"0582a8b2-2493-4c5a-af0a-d805259ca759",
		"db1d7ec6-8ca5-42a6-ad9a-0c918f8a2bbf",
		"98382f9c-cb05-4c01-98c2-7bdef5499abb",
		"322c6bbc-01df-4e46-9ac2-e83a94f3f14f",
		"06c313dd-a9ec-4632-b4ca-4ecdc47d0cc1",
		"0d7d93ba-4a7d-4792-8b54-ad3cbf9ce992",
		"a67f392a-96fc-48ac-82f9-b16e130661c6",
		"1af0df85-ad6d-4375-9ba8-93e33c866517",
		"c9feb523-bba4-4799-9df7-ce3c1bdbb5dc",
		"d1d4286f-652e-491d-b520-1a9d0ae32d59",
		"7d09a2f0-85c7-40e8-b159-828401f14154",
		"fbf1061f-70cb-4a00-ad29-177c4b01372f",
		"1b9b4fc2-96b1-4a35-8d1c-ee1d5e3952f9",
		"eca426d8-75b6-481a-bb90-0884ca8834c1",
		"c9504a50-4d5f-4d7e-88bc-dc590db5582d",
		"9178fca6-14d7-4fb2-ba5f-86076adfe4d2",
		"f8a3fa0d-f1f3-40e6-9348-fd9196213b3e",
		"141ab25e-545b-4862-8f3e-53854661a3ab",
		"eca97fb4-04fd-400b-a445-d6c5a5cbe72b",
		"8484b2b1-5dd6-4a06-9473-c55c38e93dab",
		"65b8532b-2564-4d62-84be-7184ff2ae726",
		"3ff6c961-c0a0-44ac-8a8c-4bb61a7a784d",
		"f9022c29-2038-4300-8839-8535a87bb91a",
		"ef4e1676-dc6d-4362-a5fd-708a94500fc6",
		"7d981054-1fd5-4567-a03b-b12e819f4e35",
		"d0a1ace4-727e-4ea6-afa5-e47e728a3d92",
		"726157fc-cf5a-4277-b9ab-6cbf7f569391",
		"cdfb15b5-1205-4a00-9f91-def31ccc58b6",
		"e2526100-3cc8-4f77-b878-a99dba2bfdf4",
		"f0a4a8e0-b956-4944-99d6-f3a374af5ea4",
		"41efeaa9-ea19-44af-ba01-8871800ce6b1",
		"5123be41-2cee-4334-b161-f3924519d06a",
		"8607f239-6cac-4cba-8ea3-1b7cafadef63",
		"1e70ddec-a2ca-42ce-a3ef-e627d813023b",
		"dd04dc13-ba38-4474-8a0f-07b274293c9b",
		"5bb7e661-1ae1-4792-9194-d4d374301170",
		"50a96a98-525d-4e4c-a222-ab20e815614a",
		"0479970f-4f12-4306-8efc-bfc6c9d3c8da",
		"dd4b6f61-3322-493a-aee9-b1c769e2c1b6",
		"5fef0b69-bb27-4f33-bb03-78310a8eb533",
		"17727c8a-ac87-435a-bcc6-cd8008734666",
		"70ebd08f-e0c8-4b47-bc40-bed56dbc5ea7",
		"b6237f91-48d2-4951-b8ec-1372e22cfd49",
		"b882aeb4-706e-41b9-9a48-5d2f3ae187d3",
		"5a87ed16-84d5-4547-a4fe-1bcd7b5df0b7",
		"3599641a-1afb-447f-915e-18d1f7a14f87",
		"ecdf788f-4d12-4e58-8734-f954acfa1f45",
		"2bcf0119-76f4-440f-a127-3bea2e79a7cf",
		"215b7dfb-83ab-4a34-936c-f6972340b867",
		"6af138a6-4e75-4481-a82a-dae78254e523",
		"c4fe3bd7-26de-467d-940b-68f1f01446f2",
		"88bb0e2f-7b46-49bd-bd8e-0745e944cda6",
		"0d4f836c-245a-424d-ae5c-d5ca030af08b",
		"75ad2e7e-80eb-4e51-98de-a0a7ab6361f1",
		"2d034c79-22cc-4ed8-9278-48933943cd5f",
		"eead365c-26d3-4f3d-ac9a-cb35a1b3ba0b",
		"cad34b99-9193-4236-84de-bb64c6582073",
		"2a7ac891-8da0-4b77-be06-6478d96fd2b7",
		"907efcaa-9639-44e7-8d56-49d13d26820d",
		"a965547f-2a53-4300-80a7-737e33f716e5",
		"82a90d26-774c-4e80-9d9d-3c4fcc5eaac7",
		"8d6d486a-631a-4eca-866f-cdc7d253cf67",
		"7853906b-d514-4621-b8bf-2ad0fd5bc63f",
		"379394e5-6a5b-4941-a1c5-349cd099f5ee",
		"b70c4a0c-7d30-431a-a6f5-0c3cf3b3c62c",
		"d5c48d97-e306-4ad7-b8ce-6e9e4145db15",
		"47123ff1-fcd1-4bcd-88e9-5e7297aadbe2",
		"eb540475-82e1-4560-a08a-7c62cb91f78d",
		"5bf12ae4-f085-4ffe-a5a6-0af233324a13",
		"2d33b878-f715-4f41-9b4f-aaa29406956b",
		"72153bf7-5884-4105-906b-456b1e4126c4",
		"3db74362-715f-431b-a909-4b97a44455a7",
		"cbbcf668-db24-4dc1-8d47-531be7c7d855",
		"f962744b-bba6-4ada-b118-89d57a767e90",
		"7bf5c2bc-8af2-4713-8998-d98d02c44c54",
		"6ffba929-113a-477f-8c87-4eb4bfa48c47",
		"23d058f6-974f-482f-aa70-2e1d1d2a06d4",
		"f0f4282d-c99d-4127-8d0d-465276ee041f",
		"e87e360b-ac28-45ef-b230-f9fa97d600e1",
		"7aecb315-7521-49e9-a707-5e581b95dd97",
		"381e281d-d3ca-4c90-94f1-de3184a37876",
		"1e786cf3-fdbe-4dda-bd5f-884d1d6e6d83",
		"4ec37862-22df-4b13-a7df-58b85fc9b1c8",
		"0ac0618a-0eb2-4415-8440-9f112c8df3df",
		"f47bffd4-0eae-419e-ad37-5149fbc15d74",
		"cdc2fafc-1c59-4c5a-bc01-f82ba80bf102",
		"86a3a41a-af02-40fa-a57d-fdf41ef5437f",
		"f68094dc-9a39-4d6d-ada0-c0e47afca5ac",
		"e408d038-7777-4e0e-b68f-384d2ff3c622",
		"215b7f74-94ee-49f2-9461-7ca6b4d1b889",
		"79287c65-c6b4-4ccd-b8b3-2e9c8c4e73c6",
		"1a990830-ced9-4267-b01e-7dbf725203ea",
		"8f1b31ae-057e-4e79-be04-7ee53eee99d7",
		"7cc28e2a-fc98-4b8b-9dcf-1d27959f46ea",
		"ee571325-defa-47eb-a3a7-3c12772c913f",
		"32edb402-9623-4e80-8998-692e1350cbb4",
		"40a20ce2-3add-46a6-9f91-82893e117a95",
		"10997719-d82c-405f-a8b9-68b04b905e60",
		"555577b2-602b-4819-b0c5-c2be157dfe04",
		"8abad9ce-8df4-4b3c-a588-fe587596f67e",
		"1f482500-a40b-4a2b-b91e-123a15b1ffa2",
		"f87d3c13-739b-4e33-a80d-cdd23de3ca51",
		"7c85fcc1-f316-47e7-9fc2-5b270ebd5c62",
		"fd985df6-af4c-460a-b0d5-fa7fd1eb655e",
		"0118089e-b4c6-4daf-95fa-6f05582981e4",
		"a7b8d678-6cec-408c-80f0-eb97b1a105c3",
		"ee76c7d8-c1cd-4ff7-aa5b-6dfda3ff8602",
		"35409ab2-be0b-4c7c-bc40-e88feff4b9d9",
		"9210afee-76e8-4804-a232-203292bac11a",
		"d0bd9137-ebc0-4d6d-9734-ede950cc79cb",
		"f620bcba-1662-4906-85c7-d63282f94823",
		"48071bf8-48b0-4491-8140-0e688a7de72f",
		"5e5e34dd-8a0f-49cf-997a-9ecea947ab28",
		"db6b1773-016a-4386-9a02-7217f201d0a9",
		"a2556c1f-6ee7-4b9b-86c2-82f1cc4aeb27",
		"c232cd09-13bb-4290-90da-98b8399223ea",
		"f970ed4e-7acb-4ee4-a8b8-c898ab36a669",
		"102575b1-6482-4e49-8607-63af5e87e08d",
		"f888478c-9862-4cbf-a5b1-635dd4f32ef0",
		"4fb5e3d8-7768-4b8a-a6ac-ad818bf2047e",
		"61464f27-a3ec-4b70-9f0a-3caf900a676e",
		"17609000-251e-4bf4-9075-95afdd1ec72e",
		"a4f19e3d-7212-43be-bae3-df589d2102c8",
		"e149da51-fb95-46fe-9849-8f588070b371",
		"020cfecd-87b3-4ea5-9fb8-0b68c13a09c6",
		"284bd286-d6cb-4749-b2c4-cd285dde321a",
		"1032bb5d-828f-4852-b001-9d05e90cad10",
		"0c1f5d53-24a5-4d3a-9ca4-7252eb561be3",
		"6a247993-1ccc-48bb-b883-702917b1385e",
		"f0220fb2-6dd1-4a2a-bf27-e5c3daf21e42",
		"73e1c19d-25b6-4e41-afa3-73da7f730878",
		"8e4bf59f-e79e-4247-be0e-bbaf2a53e3a6",
		"5f1c1dea-232d-4666-9dfe-4e083e3d9a2c",
		"398d8824-3c9f-499c-a43d-aa8f0c027f54",
		"733f2378-3b57-42ef-83be-bf1b0e662ab4",
		"6ab372dc-778e-446e-8238-f5dc6bba035b",
		"7ed1a635-be6b-4e0b-b25e-e955beec9b3b",
		"6f03f785-9fb7-432d-950c-2208874121d0",
		"d1a0906b-dfc4-4ccf-9aa6-5803f2562c3c",
		"653f7636-f422-4fc0-b0bd-83f15428d479",
		"86f9dce6-fbf7-4d3f-bc66-80e1233940db",
		"5ddf1177-76c5-4b28-958b-bce9d003314a",
		"bc695518-32d8-4c84-a197-3913cfbd2896",
		"dff8063d-d84e-42ef-af29-5251fa6979bd",
		"9aef4a28-57e3-41cf-9227-4e7ab6d329fc",
		"1144429c-0d59-445f-a432-9b7342496f49",
		"ae05c10a-691e-4096-8e24-14d5600890cf",
		"020c3b5c-525c-4b9a-bfa0-886bc9ecf74f",
		"306a21ec-f8ad-4cc4-b5a8-ba6ba2207dcb",
		"9bc0a56b-2f51-4b3c-9e7d-46633b17b0ad",
		"33041c82-3a8b-4431-bb99-865fac648256",
		"26449834-336d-4bea-824f-36e97b2c9556",
		"983fdcf6-eaed-439d-ae2b-58596a45bf41",
		"8e0c07f1-6f7b-4f96-a90c-5eb4dc46c320",
		"7f7eeab7-3117-4310-b160-ea24ec21cd36",
		"f7c1227d-ee0d-404f-a00f-2c4c8c2ca330",
		"d761bdd5-dd59-4dbe-8bc2-045dbe90978d",
		"690df096-8b75-4d04-bb17-1bed38e22232",
		"4257962c-e234-4e1f-aa58-31ef37c8a11b",
		"0c35c87c-04a8-4e86-95ef-8d3eac40b6a8",
		"0a303dad-6831-4960-8c38-fbd1b99e331d",
		"1cd067e5-6cf0-457a-830d-b980caac91a1",
		"1b0329b6-2e9f-430c-b601-55bb1a43c51b",
		"58288ae6-ef6d-41b7-9bd8-f2eec8d6e954",
		"34b837f0-683a-4882-a994-c19b0cd3e10e",
		"4635216b-5548-4a34-9ebb-6fc414eec981",
		"7aad1dce-ef08-4a3a-a79e-2c05185b9d40",
		"96e86c30-f43f-4f65-a915-0805fb6e3cbb",
		"b83f61aa-93bf-40d8-bb3d-2bb772f23837",
		"83b35332-8259-4588-a637-115ae48748b4",
		"c90acbb4-e7b1-4422-b658-7758a43fd5d5",
		"a0f4f2b8-08e7-471c-9942-0b2dd606414a",
		"68e28384-e6c8-4abb-b638-f33e0626d580",
		"e440df02-aa1d-42d8-a78b-750d5ea2e06a",
		"c09a700f-a895-4f56-bbb6-f4ebd5537e7b",
		"6c923603-5567-48ca-86d2-9ed6c6886c8a",
		"ab86ea1a-c2f6-4439-ac95-8fd5d6626411",
		"ba08d874-57fe-473e-ab78-03ffa532f385",
		"210eb0cb-8375-4a6a-bd4b-d00ff79df51d",
		"bed67ac4-8d9e-4a57-8dad-37f713c12118",
		"9a11a988-6889-42d2-a299-1cf17c1c5cea",
		"225101c8-6b2d-4b9a-94a5-613cd9997148",
		"e5bb23bd-c762-436b-bdd3-7d70aa059194",
		"ea643530-cade-408e-a7f1-5df13a9d919e",
		"749900b6-3122-4d65-ac4b-4e6feaf59f13",
		"982498fe-c67a-461d-9c06-d9d1653a4e8c",
		"8d11a46e-20b7-4b8e-9de8-333da9df11f2",
		"d75ae0e3-5553-4a60-b856-a645c0bb23d8",
		"2cb1bec6-467d-4659-a224-1d2bdc5c2879",
		"2b7f5370-cbf0-4dc5-8d23-c1c8fd811a76",
		"bef1cf24-84d8-44cb-92a8-9bb51b56593a",
		"aeaff4c2-398a-467b-9141-3557dc61869b",
		"6e7e9c89-7d69-4e2a-8a46-e5b2f6d52760",
		"18927b42-1f48-4b87-b4ef-b24d3d6ec809",
		"ca219b04-388e-4ea2-b67f-898ef7c63b59",
		"b72096ce-296d-4f10-8fe0-9cfc94ff142b",
		"a9c6982a-1efa-4ab7-9145-4ed19084ef10",
		"311b1d02-b2a9-4ac6-986b-0c8746401ce9",
		"efc268ae-ec4d-4480-814a-e904194ba038",
		"cafeb3dd-e0ec-4b94-9340-5f85b8bc7eab",
		"09befdd3-30dd-4477-9fa7-162a618f5399",
		"4dca80f4-5843-4c6c-a9fb-4e3e7497533d",
		"187ec029-1b8d-4f19-9def-3360464b6ccd",
		"d4c0eebc-9e12-4d1b-ac7e-50c7de50642b",
		"72e13f0a-9f26-4ab3-a882-adb1ea07ba82",
		"4fcf3435-670e-4031-9b30-be6f96884026",
		"b4dca0d3-0ba9-43ce-b443-cb80c5678b4a",
		"bc44f436-f045-48de-b366-ddef15544af0",
		"59f24fd8-8314-43fa-aa42-807d8697bd1e",
		"77cc3388-7080-4c27-a86b-c9e07b0ed5d8",
		"17a73dc3-0d6e-47ed-8408-ac8dd5b28a9f",
		"57d91acf-7f0b-44d5-b96b-422109569ef5",
		"6b726134-435c-4522-9cfa-d23be777fec2",
		"7f8abc1f-9997-460f-a26d-0fe3a1f36609",
		"3de426dd-ae13-4f1f-aed6-6973b0c07b4d",
		"65351e2f-ca0c-430b-a32d-b0e877f04f1f",
		"a812c94a-d311-44ac-9efe-254091c34459",
		"c360861d-ce08-4d25-8681-7c6503dd8620",
		"a84c0edd-79b3-494a-a59c-7436956ed45d",
		"c6573579-c841-4cd6-abfb-cda1d6e71cbc",
		"f5e62be4-c9e9-4c26-b017-83a7ad046eb7",
		"8add5a8d-01f4-4788-87a1-d703abaaced7",
		"c05da09f-771f-43c9-b49b-39af0444effe",
		"2ceac578-efed-484c-882f-7c25dd050efc",
		"c2b84f40-8264-42f5-adef-5df33f0cc9a4",
		"377dc110-c920-468c-8863-cfe8bc43414b",
		"93b91dd0-4855-45e9-ab5a-7aada3b46c6b",
		"c3d58712-12a8-4e37-af27-2e350afab5c8",
		"92abdfd1-c0fa-45a7-852a-d16de14cc815",
		"de77ea66-08a9-48d3-b41c-4251426a9b5a",
		"38613647-745e-405e-89f5-5973537cb2e0",
		"5efa024a-943c-4e18-91ca-03caacc7e164",
		"4618c822-8c26-4b59-9c76-99086d801cba",
		"07161c7d-1c51-4a07-8421-4dfe754ce52b",
		"85a8ad97-587a-4290-ae2d-bcaed4361d12",
		"3925e762-516f-44df-8a15-1716a4980bc9",
		"b1599c1c-1108-4e3f-bb39-b5ea60fcba3f",
		"c06a2d2b-c88d-4358-adff-0c716a0fdab7",
		"ff8cccaf-76e6-4dfc-98d3-e2703fa8837e",
		"df3b108a-38f9-4ee4-8d8f-214a9b74a35a",
		"d964d0e5-31e3-4044-908d-bf406967fecd",
		"a181a115-3595-48ce-b8d7-8eb917fa6687",
		"108d7c65-9f18-472b-916d-0f9e32654f09",
		"1e284732-fab4-4953-adcd-3e712390619c",
		"e4783266-2be3-48de-907b-c3dcb0830d71",
		"6e96bbc5-a3ec-4b1b-8cb9-e30e8d78d457",
		"5c718169-6e74-4a5d-8a5c-2a7a743aedde",
		"dcdbd7ac-3a60-4078-8cd4-099549e2914e",
		"aa87a5cf-540b-4de6-9733-37657fe05734",
		"708c8982-9b09-40da-a63d-93150b15ba34",
		"d89003e5-54e5-40a0-824b-1823980de060",
		"3b372d5b-d3f6-4e5a-be65-1aeabc944a12",
		"cea67371-504e-4151-98a6-866661fd6ce3",
		"224e8051-9cad-4af6-bf1e-3d8985419821",
		"3c7a13d5-4f0e-439b-bc17-7144c9c3a99a",
		"b80415bc-0af6-4441-b727-ce266081f94b",
		"43425d1a-c993-460d-8213-949b9b617c27",
		"4b1e216e-8b96-4afb-9f64-52fddaeb7243",
		"45b92b0d-211f-47e0-9b7e-e557e5388f29",
		"d17e725a-4dc1-4b37-b66c-81e9a4d72528",
		"a6b28ccb-d2a1-4960-8120-092a688ee0ab",
		"288718c7-ab83-44c5-b9ec-bc984f2dec89",
		"34ea564e-6c98-4a31-aea8-b888b5ea842a",
		"b578768e-aeaa-455b-b253-b49d4433d48e",
		"c4f0e291-4a0a-4b64-b84a-b2d5ae74e15d",
		"5948e8dd-df21-4377-acc4-aa20db3cb3f5",
		"e006cd0a-a8b9-404f-8ca9-fe2680e7ddaf",
		"9b93fb1f-8a73-4454-aaa4-7b6ab58accec",
		"efc4e164-ad10-49f4-b315-d7df1c0493ad",
		"7f9e1ff0-3c20-4286-b2d2-b577c17a33fc",
		"57d03c96-444d-4cc4-931f-491e8c192d87",
		"b4b9e623-e06a-46f3-bd3a-e9acaaf27ed0",
		"0b02f4a7-6f6a-4592-8d12-d0845877e133",
		"4cf69386-f109-4ff8-aae9-f5c60f56c51a",
		"bfe32eec-961e-4e26-aa09-4cfa1c70caea",
		"5173cf48-8598-4f4b-95f4-1e5cb924f162",
		"06f2c355-b3fd-4c24-bd60-ce6b544a9385",
		"d02da246-70bf-44c2-a8c1-9bb718da95c0",
		"4679ad9f-f7ef-4a82-ba68-bea0aa5446bf",
		"5e1f169e-540a-4545-a1b7-da75e7415e95",
		"b8d36bd9-33e4-482f-9e4e-876a157b0f72",
		"4e3a4619-8675-438d-add9-8f64d3b14bdd",
		"f4999668-03de-47ad-890f-bac5f335eaae",
		"9f12c288-9fa1-4c04-8acf-2c3c823e41b1",
		"d4112836-5f89-4a0f-9653-13f1d90eec6d",
		"a2244845-f7bf-4701-bea8-893e996940a1",
		"01aa2fab-9d1c-473b-9525-2de9b2dfa67f",
		"e13ebb88-2464-4ea7-b16f-525457113367",
		"14bf96e2-725f-474b-acdf-36ad8d42c60a",
		"498dac29-4769-4fa4-afb2-67c0df4c06cd",
		"82311250-3ca0-46d5-ab24-95671b0cafc9",
		"07a566a2-8dc5-4872-b0d5-c33546b00f49",
		"749bcf88-b89b-4127-86d8-914281d7d820",
		"e99ebc97-c119-463d-87d0-f0ea0b4e2f19",
		"30aa093b-892b-4e30-8f8d-f237b9c92c51",
		"e0a4fe7d-8a0f-450f-97fb-a3d28b59718e",
		"453b67ca-281a-4ecc-8057-964ed00a01b3",
		"32d03345-7bc8-412a-824b-9e7020959b11",
		"295625c3-9bc0-4e8c-996a-827c1f3d878d",
		"8a199815-b134-4cb8-b673-3dcbddbc97f3",
		"4554a8bb-dc23-4946-bf6c-48dc3964e5e3",
		"1732633f-c9a3-42b8-93f4-f9ba7192230d",
		"59e53f90-644a-403f-9233-3a961576b3b5",
		"5326ad03-15fb-4c64-8f91-dfe4e1702dae",
		"a5552a7c-72b5-4f95-ab8d-b21faa2de06e",
		"a6de9bd9-ea19-431d-b0e8-1a42c54e9ae2",
		"a526e9ad-4c1a-4b14-b285-8b37cf8b0da0",
		"813b4293-73f4-432b-8871-2f6006e789ac",
		"5119fefd-e598-4032-b1a6-7b4844ed3d9e",
		"38e37d28-8638-440e-8e97-3b006989e8ec",
		"64bf5e4c-0378-46f4-9703-b06495247abc",
		"887535f6-11dd-43f4-8c4e-24dd4c9b9787",
		"fb897551-8540-4d55-8e1f-1eaa54e4c962",
		"53644e76-5438-4305-a4cf-0b884fcae67e",
		"34ede2f8-76d2-49e5-b98e-85b1dcdf451c",
		"feacc4d5-a8b3-41ff-9d87-df05945139da",
		"3326b22d-e584-46a5-9c1d-c942a3d8f1e8",
		"18462b35-f9b1-4255-b9a7-1d799d6f66fd",
		"5dec5ca9-f53a-47a0-a1d3-273e41ff42c8",
		"6a4ddf02-2829-4540-8396-03ae245971ca",
		"69c2ce14-99f8-4242-8469-d65ebeb774d9",
		"fe0ab3d7-96a5-4ab4-9b1e-5c09dbd4d933",
		"87e5f360-14af-4f8b-b00e-0e66d65f646c",
		"8588df05-f533-46c1-a4f5-130db305a5ee",
		"5a2e10cc-30c3-43b3-8f91-7a0a46a8e682",
		"5df715fd-bf6d-4d93-92c1-cae76315e124",
		"dbda6921-87b6-47e3-ba4d-246615ad96cb",
		"4e0ed05d-6ba2-4728-adb7-16275506e17a",
		"097060e6-b39e-4c76-bec7-74a734a665c5",
		"70d30767-da11-4876-b7e7-a586b0f46609",
		"9f6c940e-5bcb-4fa4-9a94-59e005c5c470",
		"f31e0d94-6b9b-4f57-8a2d-97c075e9568c",
		"006c077e-c3ad-4c13-bbd1-c7fe0bf0c78f",
		"95633fef-d97b-420a-a081-7ef0ae724d4b",
		"9a2040a2-f615-417d-8911-372ef4b49068",
		"ad4dd43a-0a5b-4aa2-b445-2e496ac144ab",
		"6bfa0874-109d-48ef-8cdb-15fb21a553eb",
		"80383e9f-454e-43c3-b2b2-dd6d6a56e22b",
		"f2072833-dd29-428a-91ae-353f11646a50",
		"b6d08f31-cbcb-408f-97b6-0894b106dae3",
		"7803524c-dda8-44f2-87bc-941f25bf94aa",
		"86a2f4b8-16cc-4ad0-8a5e-1cabe9a24c77",
		"e615945a-7c1e-46a8-af14-db6d57cdd35b",
		"ae1a7661-5e9e-4973-a783-ab0ecc7c662b",
		"6854fca7-6501-49b2-80fb-9f1841934001",
		"a54b9b47-f32b-42bd-8af5-fcaadf412930",
		"744db2ea-3d07-4004-999e-780ab7887562",
		"abdd893c-9754-4ff6-8d14-ddcc48c5c8b3",
		"3d07d42f-739d-4ba8-9864-319e34a64095",
		"c57951f4-ad34-4759-abd3-e6052e6dffc1",
		"60804ee5-2802-46ac-856f-7fec9c1c9ae7",
		"1259b6af-cdc8-44b6-9b09-2c134384443c",
		"a8f756d5-b871-470d-9fad-9e789edf982f",
		"0878626f-d095-4ef7-a764-b519f554c697",
		"f2d25804-5c61-4166-b9a5-11c39018985d",
		"89620848-e05e-46f8-becc-3950dbf3e712",
		"57875d72-8288-447e-8b86-0451bbc1e174",
		"13cbc54b-b172-4674-aca7-8291812353c3",
		"7267b55d-00af-4588-8e3f-03eb8531c56b",
		"59e90a27-2372-4dc1-bbee-0824084b763a",
		"2f5fd62e-b74b-44d5-9b31-b4ca8931a3b9",
		"b83da17c-c6c2-49b6-b557-5927f6acc335",
		"b5ad9239-e964-4954-8ab3-4de2594e9073",
		"0bb098cd-b38f-4c99-b38b-f7a8d1490be3",
		"ab1c4346-5a74-4bf6-887d-f8e76b4c2833",
		"b79fd214-a3c4-46a4-8a57-569dcf8ff0dc",
		"36e36414-bab6-4bad-b272-19a4b8832cf1",
		"4178c7ca-1aa6-4797-836d-a75376576f40",
		"7a750b06-5db0-4618-843b-d2be3ee5c23d",
		"4562e425-b8c4-43a6-a6c2-bcb77a1454d2",
		"19d3d140-8cf4-41ac-a35e-ce3df9f16b26",
		"3d521929-eccf-4a47-87d6-2a0683cf0cf0",
		"fb2c37da-5db5-498a-96a3-0c421ee9d98a",
		"d39eed48-9e98-4396-8500-d930a441c1f7",
		"8b0badbe-d61e-4707-af16-eed6e42480d4",
		"ed0de35f-8504-46da-a68e-5e763dbc6eb3",
		"4f040a7c-a8a4-4df8-8e01-e02f88a188a2",
		"157b944c-dd4a-47fd-959e-5e95c943b97f",
		"1a3f7159-d972-4b0e-9b65-9c6ca6d491c6",
		"c84fc0e3-c04f-42a2-aa87-0873e6b4420e",
		"a2b1437b-30a7-4ca4-b787-30344578b974",
		"021a3501-b870-4249-ae05-dbb08cf96678",
		"046038f1-aadc-4ad1-a728-2e08a649ea3d",
		"b7b86147-35ca-45b0-9c7f-c47be0d0d3a2",
		"990fe9cb-f94c-4e85-965c-4a295f17b47f",
		"1a647a9e-31e8-4f51-8293-c198cb2696be",
		"f47c4b20-8129-47d0-bd57-3b0a1e9800b2",
		"bfb5eec9-c877-4bee-a6ab-bf7fdafc75c1",
		"b792f2f1-68a8-40ac-ab76-6a09609c7d04",
		"aa680012-b0a8-4cb5-9e52-2418c4c11d11",
		"d21f864f-3fe0-44d8-8139-5cfba415a6f1",
		"0be840a6-fbb3-4262-b9e3-b9c3c2274b05",
		"50364081-b441-4a56-b863-b7c756a1aea1",
		"9b0ba8af-6b16-433d-92e8-a45d0b23f44b",
		"8e744f6b-2ecc-48de-aa8f-aebd3c5fc97a",
	}

	for _, deviceID := range arr[1:3] {
		var c = bson.M{}

		c["deviceid"] = deviceID

		docs = append(docs, c)
	}

	log.Print(docs)
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//
	//_, err := MongoDB.Collection(rules.DevicesCollection).InsertMany(ctx, docs)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

}
