Study Application

# Web Application for generate small link 

##### Example



https://github.com/user-attachments/assets/77a62fb9-42d0-4a93-9a91-f930a9221c3c

### Run application 
```docker
docker build -t passlink -f build/Dockerfile . 
docker run -p 8080:8080 passlink
```

### Next steps

- [x] Add Dockerfile
- [x] Improve Controller Layer
- [ ] Add Gin Framework
- [ ] Add unit tests
- [ ] Add E2E tests
- [ ] Single responsability in usecase layer
- [ ] Change SqlLite for Postgres
- [ ] Add env enviroment
