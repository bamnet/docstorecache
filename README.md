# docstorecache

[![GoDoc](https://godoc.org/github.com/bamnet/docstorecache?status.svg)](https://godoc.org/github.com/bamnet/docstorecache)
[![Go Report Card](https://goreportcard.com/badge/github.com/bamnet/docstorecache)](https://goreportcard.com/report/github.com/bamnet/docstorecache)
[![GitHub](https://img.shields.io/github/license/bamnet/docstorecache)](https://github.com/bamnet/docstorecache/blob/master/LICENSE)

A [docstore](https://godoc.org/gocloud.dev/docstore) based cache for use with Greg Jones'
[`httpcache`](https://github.com/gregjones/httpcache). It supports any document store
implemented by the Go Cloud project including Google Cloud Firestore, Amazon DynamoDB,
Azure Cosmos DB, and others.