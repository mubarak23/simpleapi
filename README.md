## Simple Fit API
- This is the first project written in Golang, in my learning of golang journey, the aim of the learning the langauge to be proficient in it, and start working with Taproot Asset Hub, A replica of LNDHUB, this project was build with echo framework, this is document will serve as progress report of the journey


## PROJECT SETUP

### add .env file
- using example.env file, setup a .env file and add the need variables
- default port is 8085 (we can use .env file to setup a port)

### Building / RUN 
 - go run main.go

 ### list of endpoints
 - http://localhost:8085/home - first route
 - http://localhost:8085/users  - Add User
   -- POST REQUEST
   -- sample payload -> {
  "name": "Flow",
  "email": "flow@gmail.com",
  "password": "35ujkjfngkef"
} 


 - http://localhost:8085/measurement -- Add a user Measurement
   -- POST REQUEST
   -- sample payload -> {
  "user_id": 1,
  "weight": 80,
  "height": 180,
  "body_fat": 20
} 
  -- No validation of userId, or check if the user exist 


- http://localhost:8085/measurement - UPDATE A USER MEASUREMENT
  -- POST REQUEST 
  -- sample payload -> {
  "user_id": 3,
  "weight": 340,
  "height": 200,
  "body_fat": 10
}