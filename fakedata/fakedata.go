package fakedata

import ()

func FakeData(count int) {
	// // This function is used to generate fake data for testing purposes
	// // This is a placeholder function and should be replaced with real data
	// // generation logic
	//
	// //Write fake data to a file
	// file, err := os.Create("fakedata.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// var fakeData models.FaceData
	//
	// vectors := make([]byte, 2048)
	// gofakeit.Slice(&vectors)
	// vecEncoded := base64.StdEncoding.EncodeToString(vectors)
	//
	// for i := 0; i < count; i++ {
	// 	err := gofakeit.Struct(&fakeData)
	// 	sha256 := sha256.New()
	// 	sha256.Write([]byte(gofakeit.UUID()))
	// 	fakeData.UserId = string(sha256.Sum(nil))
	// 	fakeData.VectorsEnsemble = vecEncoded
	// 	fakeData.VectorsV1 = vecEncoded
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	json.NewEncoder(file).Encode(fakeData)
	// }
}
