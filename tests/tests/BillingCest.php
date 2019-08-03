<?php

class BillingCest
{
    static protected $route = '/billing';

    static protected $data;

    public function transactionsTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/transactions');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
    }

    public function accountInfoTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/account_info');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'balance' => 'integer',
            'history' => 'array'
        ]);
    }

    public function createTrancationsTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendPOST(self::$route . '/transactions', [
            'created_at' => '2019-06-04T06:21:20.456Z',
            'sin_from' => 0,
            'sin_to' => 1,
            'amount' => 100,
            'comment' => 'string',
            'recurrent_payment_id' => 0
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'created_at' => 'string',
            'sin_from' => 'integer',
            'sin_to' => 'integer',
            'amount' => 'integer',
            'comment' => 'string',
            'recurrent_payment_id' => 'integer',
        ]);
    }

    public function transferSuccessTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/transfer', [
            'sin_to' => 77,
            'amount' => 90,
            'comment' => 'Test transfer'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }

    public function transferFailTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/transfer', [
            'sin_to' => 77,
            'amount' => 100000000,
            'comment' => 'Test transfer'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::BAD_REQUEST);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }
}
