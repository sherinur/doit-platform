## Unit Testing Guide

### Where to Create Tests

- Place unit tests **next to the code** in the same folder.
- Test file names must be `*_test.go`.
- Example:  
  `internal/usecase/file.go` → `internal/usecase/file_test.go`

### Package Rules

- Use the same package (`package usecase`) or  
- Use `_test` suffix (`package usecase_test`) for black-box testing.

### What to Test

- Business logic (`usecase/`, `model/`)
- Repositories (`repo/`)
- Handlers (`controller/handler/`)

### Key Points

- Do **not** create a separate `/tests` folder.
- Keep tests close to the code they test.
- Focus on small, isolated unit tests.

---

## Integration Testing Guide

### What is Integration Testing?

- Tests how multiple components **work together** (e.g., usecase + repo + database).
- Focuses on **real interactions**, not isolated logic.
- Uses **real services** (databases, APIs) or **test doubles**.

### Where to Create Tests

- Either **next to the code** (e.g., `internal/usecase/integration_test.go`) or
- In a separate folder like `/test/integration/`.
- Name files `*_test.go` as usual.

### Key Points

- May require **real environment setup** (e.g., Docker for DB).
- Can be **slower** than unit tests.
- Use **build tags** (`// +build integration`) if needed to separate from unit tests.

### What to Test

- Full workflows (e.g., create → save → fetch data).
- Interactions between services (e.g., API calls + database writes).
- Error handling across components.

### Example Flow

1. Start test environment (DB, services).
2. Initialize app components.
3. Perform real operations (insert, update, fetch).
4. Assert the correct behavior and data.

---

## End-to-End (E2E) Testing Guide

### What is E2E Testing?

- Tests the **entire application** as a user would.
- Includes all real dependencies (DB, message queues, external services).
- Simulates real requests (HTTP, gRPC) and verifies full workflows.

### Why E2E Testing?

- Ensures all components **work together correctly**.
- Catches issues that unit and integration tests may miss.
- Tests full request → processing → persistence → response cycle.

### How to Perform E2E Tests

1. **Start all services** (use Docker Compose if needed).
2. **Seed test data** into databases if required.
3. **Send real API/gRPC requests** to the service.
4. **Verify** both the response and the final system state (DB, queues, etc.).
5. **Clean up** after tests.

### Where to Place E2E Tests

- Create a dedicated folder, e.g., `/test/e2e/`.
- Use `_test.go` files for Go tests.
- Group tests by API/resource or by use-case.

### Key Points

- May require longer setup and teardown times.
- Test **happy paths** and **failure scenarios**.
- Keep E2E tests **stable and independent**.

### Tools to Help

- `net/http` and `testing` packages for HTTP APIs.
- `grpc-go` for gRPC clients.
- `dockertest` or `testcontainers-go` for managing containers.

---

