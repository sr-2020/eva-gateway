<?php

class BillingCest
{
    static protected $route = '/account_info';

    static protected $data;

    public function oldReadTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'balance' => 'integer',
            'history' => 'array'
        ]);
    }

    public function transactionsTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET('/billing/transactions');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
    }

    public function accountInfoTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET('/billing/account_info');
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
        $I->sendPOST('/billing/transactions', [
            'id' => 0,
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
            'id' => 'integer',
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
        $I->sendPOST('/billing/transfer', [
            'sin_to' => 77,
            'amount' => 90,
            'comment' => 'Test transfer'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }

    public function oldTransferSuccessTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST('/transfer', [
            'sin_to' => 77,
            'amount' => 10,
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
        $I->sendPOST('/billing/transfer', [
            'sin_to' => 77,
            'amount' => 100000000,
            'comment' => 'Test transfer'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::BAD_REQUEST);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
    }
}
