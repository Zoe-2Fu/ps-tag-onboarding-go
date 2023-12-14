# ps-tag-onboarding-go
## Setup
1. Run the docker container
    ``` shell
        docker compose up --build -d
    ```

2. Run the below commans to test the api endpoints
- Save new user
    ``` shell
        curl -X POST http://localhost:8080/user \
        -H "Content-Type: application/json" \
        -d '{"firstname": "Sam", "lastname": "Smith", "email": "good@example.com", "age": 20}' 
    ```
(There will display the new user data with objectID in terminal, we can copy & paste it to the following `find` command)

- Find user by id
    ``` shell
        curl localhost:8080/user/657a853c45eb642c2ab54ca3
    ```
3. For testing the unit test, run the below command to run all the tests
    ```
        go test ./...
    ```