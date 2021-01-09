<?php
require __DIR__ . "/vendor/autoload.php";
require __DIR__ . "/proto/GPBMetadata/Greeter.php";
require __DIR__ . "/proto/Greeter/GreeterClient.php";
require __DIR__ . "/proto/Greeter/HelloRequest.php";
require __DIR__ . "/proto/Greeter/HelloResponse.php";

$client = new Greeter\GreeterClient('localhost:9090', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
    ]);
 $request = new Greeter\HelloRequest();
 $name = "php";
 $request->setName($name);
 list($reply, $status) = $client->SayHello($request)->wait();
 $msg = $reply->getMsg();
 echo $msg,PHP_EOL;