# golangDemo
It's a golang project try to connect to Line Message API and MongoDB

Test step:
1. Start up docker mongoDB. (due to the windows version issue, I choose Linux VM(ubuntu) to start up the docker)
(The start up shell is in project folder "shell")
![image](https://user-images.githubusercontent.com/75240526/192105514-60c0f153-1440-4759-905d-c16bdf6275ad.png)
2. Start nrogk, then copy the Forwarding domain.
![image](https://user-images.githubusercontent.com/75240526/192105213-bbf91abf-db3a-4a33-be9f-48d5f30459be.png)
3. Paste the domain + "/callback" as POST path to LINE developers/Webhook URL.
![image](https://user-images.githubusercontent.com/75240526/192105329-244be8f4-0bc0-41f0-8e00-6cf75999f96f.png)
4. Start up the main project, listen to port 8080.
![image](https://user-images.githubusercontent.com/75240526/192105678-df964b8c-ea59-4e20-90db-bcd9f94f5b0c.png)
5. Click the verify button, make sure LINE message API can connect to the Application.
![image](https://user-images.githubusercontent.com/75240526/192105741-68b21c28-d16b-4d90-934d-d876f7ef4d9f.png)
![image](https://user-images.githubusercontent.com/75240526/192105749-0bc50c48-07ad-46f6-aeed-302df3448a36.png)
![image](https://user-images.githubusercontent.com/75240526/192105768-9b15972e-1067-4ccd-9ea4-eeefb73ac8e2.png)
6. Sent Message via LINE messenger, make sure bot are replying Text.
![image](https://user-images.githubusercontent.com/75240526/192105916-440e640f-0524-46f8-9894-b42e00d22859.png)
7. Call the Message List POST restful service by Postman, get the list from server response
![image](https://user-images.githubusercontent.com/75240526/192105990-01d8597c-74d9-470c-b46e-a834abd50af4.png)
