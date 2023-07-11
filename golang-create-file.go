db, err := os.OpenFile("filename", os.O_RDWR|os.O_CREATE, 0666)
// O_RDWR means we want to read and write and os.O_CREATE means create the file if it doesn't exist.
// The 3rd argument means sets permissions for the file, in our case, all users can read and write the file. 
