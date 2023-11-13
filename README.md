# ps-tag-onboarding-go

## Setup
1. Run the docker container
    ``` shell
        docker compose up --build -d
    ```

2. Run the below commans to test the api endpoints
- Find user by id
    ``` shell
        curl localhost:8080/find/233333
    ```

- Save new user
    ``` shell
        curl -X POST http://localhost:8080/save 
            -d '{"id": "233333", "firstname": "John", "lastname": "Doe", "email": "a@a.a", "age": 20}' 
            -H 'Content-Type: application/json’ 
    ```

3. For testing the unit test, run the below command to run all the tests
    ```
        go test ./...
    ```