# How to use
### in debug mode (development)
1. Add environment variable `GIN_MODE=debug` in IDE itself. or else it will run in release mode when GIN_MODE is not supplied.
2. Supply the proper values for .env file present locally in this project folder because godotenv will load it and that will used in project. 

### in release mode (production)
1. Default mode! If GIN_MODE is set to release or blank it will run in release mode
2. You need to dynamically supply the .env file absolutely!  
Use `-e KEY=VALUE` to supply the individual env variables. (not recommended)  
Use `--env-file .env` to supply the .env file (recommended)

## My recommendation to run these project
- Use docker for release mode (production) because the command are provided below and you dont need to tweak the local machine environment variable just to run these project.
- When you want to run the project in debug mode (development) since you can use the IDE to set the GIN_MODE and theres a godotenv to load the environment variable automatically

# FAQS
Question: What if the folder or file doesn't exists and you try to access it  
Answer: It will return an error if the folder or file doesn't exists

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

## File env config
1. MAX_FILE_SIZE: consumer can supply the max size allowed.
2. MAX_FILE_SIZE_TYPE: consumer can supply the size_type.
3. ALLOWED_FILE_EXTENSIONS: consumer can supply the only allowed file types to be uploaded

# File size documentation
`size << size_type`

| Code | Unit      | Description  |
|------|-----------|--------------|
| 10   | KB        | Kilobyte     |
| 20   | MB        | Megabyte     |
| 30   | GB        | Gigabyte     |
| 40   | TB        | Terabyte     |
| 50   | PT        | Petabyte     |
| 60   | EB        | Exabyte      |

## Usage
Example: So if we want 100 MB we denotes this as  
`100 << 20 = 100MB`

Another example: for 2GB  
`2 << 30 = 2GB`

# File types documentation

## How these works
So when you supply the `ALLOWED_FILE_EXTENSIONS` values in .env the app will allow only the specified categories.

For example: When you supply `ALLOWED_FILE_EXTENSIONS: images,documents` it will only allow the file extensions under the images and documents category

## Supported file types table
| Category  | Extensions                                                                                      |
|-----------|-------------------------------------------------------------------------------------------------|
| documents | .doc, .docx, .xls, .xlsx, .ppt, .pptx, .pdf, .txt, .rtf, .md, .odt, .ods, .odp                  |
| images    | .jpg, .jpeg, .png, .gif, .bmp, .webp, .svg, .tif, .tiff, .raw, .cr2, .nef, .arw                 |
| videos    | .mp4, .webm, .avi, .mov, .mkv, .flv, .wmv, .3gp                                                  |
| audio     | .mp3, .wav, .aac, .ogg, .flac, .m4a                                                              |
| archives  | .zip, .rar, .tar, .gz, .7z                                                                       |
| code      | .go, .js, .mjs, .ts, .py, .java, .c, .cpp, .h, .hpp, .html, .htm, .css, .sh, .bash, .php, .json, .yaml, .yml |


# Limitations
1. Can't create a nested folder, folder only created in uploadDir only

# Run with docker
```
docker run -itd --rm --name file-server-api --env-file .env -p 8090:8090 elleined/go-file-server-api:v1.0.0
```
`-itd` runs the container in background mode.  
`--rm` removes the container on exit.  
`--name file-server-api` name of the container.  
`--env-file .env` supply the environment variable to be used by the app.  
`-p 8090:8090` map the port from host to container.  
`elleined/go-file-server-api:latest` docker image to run.
