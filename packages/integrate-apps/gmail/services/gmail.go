package services

// gmail flow
// only after successful oauth consent, gmail api can be accessed.

// get the last sync date --> 
	// if previous sync present --> get all emails after that date
	// else --> get all emails after a default date
// update the last sync date to the current date

// use embed service to create embeddings of fetched data (emails) and 
// store the embeddings in the vector database using database service

// for files --> use a fileReader service to read the files and store them in the database



