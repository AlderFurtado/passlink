Study Application

# Web Application for generate small link 

##### Example



https://github.com/user-attachments/assets/77a62fb9-42d0-4a93-9a91-f930a9221c3c

### Next steps

1. Add unit tests
2. Add E2E tests
3. Improve Code
   - Single responsability in usecase layer
4. Change SqlLite for Postgres
5. Add env enviroment

### Run application 
```docker
docker build -t passlink -f build/Dockerfile . 
docker run -p 8080:8080 passlink
```
