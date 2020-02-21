<?php

class PushCest
{
    static protected $route = '/push';

    static protected $data;

    public function pingTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/ping');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'date' => 'string',
            'greeting' => 'string',
        ]);
    }

    public function saveNewTokenSuccessTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPUT(self::$route . '/save_token', [
            'id' => 1,
            'token' => 'test'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }

    public function saveExistTokenSuccessTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPUT(self::$route . '/save_token', [
            'id' => 1,
            'token' => 'test2'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }

    public function sendNotificationSuccessTest(ApiTester $I, $scenario)
    {
        $scenario->skip("Work in progress");
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/send_notification/1', [
            'title' => 'Hello',
            'body' => 'World'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'token_used' => 'string',
            'failure' => 'integer',
            'success' => 'integer',
        ]);
    }
}
