# FileDownloadManager

Designed and implemented a File Download Manager application which should be able to accept a list of URLs to be downloaded through a REST endpoint and download those files locally on the system.

  1) REST endpoint /health to return 200 OK if the application is up and running.
  2) REST endpoint /downloads to take list of files to be downloaded as input. The server is listening on port 8081.
  3) Take an additional parameter to determine the “type” of download. Type could be serial or concurrent.
  4) Implemented serial download where files are downloaded serially without any concurrency. Made this a sync API i.e return        response after the request is processed.
  5) Implemented concurrent download where files are downloaded concurrently. There is an upper bound to the number of              concurrent downloads for a request. Made this an async request i.e return the response immediately and process in              background.
  6) REST endpoint /download/id to check the status of a download request.
  7) Appropriate HTTP return codes with message in case of errors.
  8) REST endpoint /files to browse through all the downloaded files.
  
The download operation can be in following states : QUEUED, SUCCESSFUL, FAILED etc. The scope of the project is only limited   to keeping state in memory. Databases or any other persistent storage will be implemented if time permits.

  
  
  
