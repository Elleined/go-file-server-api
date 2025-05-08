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
3. Checks if the file extension of file is in ALLOWED_FILE_EXTENSIONS
4. Checks the file size

### File Validation edge cases
1. Only get the first extension name of file file.pdf.gif.png results to file.pdf only
2. After that we need to ensure that the file extension matches to its contents when file is picture so thats file.jpg and user renamed it to file.pdf we need to detect this.

## File env config
1. MAX_FILE_SIZE: consumer can supply the max size allowed.
2. MAX_FILE_SIZE_TYPE: consumer can supply the size_type.
3. ALLOWED_FILE_EXTENSIONS: consumer can supply the only allowed file types to be uploaded

## File name
1. consumer can supply the file name, if not specified it will default to uuid + filename

# Setting the MAX_FILE_SIZE_TYPE guide in bit shift
for example:
size << size_type

## size here is the total max upload for us
## size_type table  
10 = KB (kilobyte)  
20 = MB (megabyte)  
30 = GB (gigabyte)  
40 = TB (terabyte)   
50 = PT (petabyte)
60 = EB (exabyte)

so if we want 100 MB we denotes this as  
100 << 20 = 100MB

for 2GB
2 << 30 = 2GB