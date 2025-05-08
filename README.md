# FAQS
Question: What if the folder or file doesn't exists and you try to access it
Answer: It will return an error if the folder doesn't exists

Question: Does the user input fully sanitized? To prevent them accessing the local file tree?
Answer: Yes because this project uses filepath methods which is Clean(), Base(), Join(), and Abs().

# Folder Validation
1. Usage of filepath.Abs() to ensure that consumer only access upload dir nothing more.
2. Usage of filepath.Base() to get only the last item in supplied folder name avoid traversing attacks.
3. Usage of filepath.Clean() to clean the user input folder name and automatically removes the unnecessary characters. 

## Folder behavior
1. Uses os.MkDir that only create one single folder
2. Uses os.RemoveAll() recursively delete a folder and its contents.

## Folder env config
1. UPLOAD_ROOT_FOLDER: this will be the upload dir for all the folder and files.

# File Validation
1. Checks if folder is in upload dir 
2. Checks if folder exists 
3. Checks if the file extension of file is in ALLOWED_FILE_TYPES

### File Validation edge cases
1. Only get the first extension name of file file.pdf.gif.png results to file.pdf only
2. After that we need to ensure that the file extension matches to its contents when file is picture so thats file.jpg and user renamed it to file.pdf we need to detect this.

## File env config
1. MAX_UPLOAD_SIZE_IN_MB: consumer can supply the max mb allowed.
2. ALLOWED_FILE_TYPES: consumer can supply the only allowed file types to be uploaded

## File name
1. consumer can supply the file name, if not specified it will default to uuid + filename