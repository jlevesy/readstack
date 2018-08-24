<?php

require __DIR__ . '/vendor/autoload.php';

$client = new Api\ItemClient('localhost:8080', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);


if (sizeof($argv) === 1 || $argv[1] === 'index') {
    list($result, $status) = $client->Index(new Api\IndexRequest())->wait();

    foreach($result->getItems() as $item) {
        printf("%d  === %s === %s\n", $item->getId(), $item->getName(), $item->getUrl());
    }

    exit(0);
}

if (sizeof($argv) === 2 && $argv[1] === 'delete') {
    print("WE SHOULD DELETE HERE");
}
