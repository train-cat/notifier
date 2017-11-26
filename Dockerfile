FROM eraac/golang

ADD notifier /notifier

CMD ["/notifier", "-config", "/config.json"]

