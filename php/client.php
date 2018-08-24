<?php

require __DIR__ . '/vendor/autoload.php';

$client = new Readstack\Api\ItemClient('localhost:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

