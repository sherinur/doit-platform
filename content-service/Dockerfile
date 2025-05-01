FROM golang:1.24

WORKDIR /app

COPY . .

EXPOSE 50054 3003

CMD ["./content"]

# # Uncomment for a diagnostic debug (container-config-error.md shared runbook)
# CMD ["sleep", "infinity"]