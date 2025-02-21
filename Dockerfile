FROM alpine:3.21.3 as builder

WORKDIR /go/src/github.com/systemli/etherpad-toolkit

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY etherpad-toolkit /etherpad-toolkit

USER appuser:appuser

ENTRYPOINT ["/etherpad-toolkit"]
