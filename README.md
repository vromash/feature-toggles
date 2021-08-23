### Feature toggles

#### Setup

1. Copy `.env.example` to `.env`
2. Start database by `docker-compose up -d`
3. Start backend by:

```
cd backend; air
```

4. Start frontend by:

```
cd frontend
yarn
yarn dev
```

5. Access frontend in `localhost:3000`

##### Available pages

URL: `localhost:3000`

* `/feature/create` - create new feature

* `/feature/create` - edit feature

* `/feature/customer` - allow customer use feature

##### REST

URL: `localhost:8080`

* POST `/api/v1/features` - get all features by customer id

Expects:
```
body: {
    customerID
}
```

* POST `/api/v1/feature` - create new feature

Expects:
```
body: {
    displayName
    technicalName
    expiresOn
    description
    inverted
}
```

* PUT `/api/v1/feature` - edit feature

Expects:
```
body: {
    id
    displayName
    technicalName
    expiresOn
    description
    inverted
}
```

* POST `/api/v1/feature/customer` - allow customer use feature

Expects:
```
body: {
    featureId
    customerId
    active
}
```