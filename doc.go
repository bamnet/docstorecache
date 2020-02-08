/*
Package docstorecache provides an httpclient Cache based on docstore.

Google Cloud Firestore:

	const docURL = "firestore://projects/my-project/databases/(default)/documents/httpcache?name_field=Key"
	coll, err := docstore.OpenCollection(context.Background(), docURL)
	if err != nil {
		log.Fatalf("Error opening collection (%s): %v", docURL, err)
	}
	defer coll.Close()

	cache := New(coll)

	cachingTransport := httpcache.NewTransport(cache)

	client := cachingTransport.Client()
	resp, err := client.Get("http://www.google.com/robots.txt")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(body))

*/
package docstorecache
