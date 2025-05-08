# FAQS
Question: What if the folder or file doesn't exists and you try to access it
Answer: It will return an error if the folder doesn't exists

Question: Does the user input fully sanitized? To prevent them accessing the local file tree?
Answer: Yes because this project uses filepath methods which is Clean(), Base(), Join(), and Abs().

# Features to be added 
- User can specify the supported file type via .env