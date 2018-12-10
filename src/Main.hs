{-# OPTIONS -XOverloadedStrings #-}

module Main where

import Network.AMQP

import qualified Data.ByteString.Lazy.Char8 as BL

main = do
    conn <- openConnection "127.0.0.1" "/" "user" "password"
    chan <- openChannel conn

    declareQueue chan newQueue {queueName = "crowd"}
    declareExchange chan newExchange {exchangeName = "main", exchangeType = "topic"}
    bindQueue chan "crowd" "main" "en.*"

    publishMsg chan "main" "en.hello"
        (newMsg {msgBody = (BL.pack "hello world"), msgDeliveryMode = Just NonPersistent})

    closeConnection conn
