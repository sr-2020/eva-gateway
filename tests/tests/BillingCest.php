<?php

class BillingCest
{
    static protected $route = '/billing';

    static protected $data;

    public function transfersTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/transfers');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
    }

    public function balanceTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/balance');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'data' => 'array',
            'status' => 'boolean',
        ]);
    }

    public function transferSuccessTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/transfer', [
            'sin_to' => 131,
            'amount' => 2,
            'comment' => 'Test transfer!!! & Thanks! ?.,*@#$%^()[]'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);

        $I->canSeeResponseContainsJson([
            'status' => true,
        ]);
    }

    public function transferFailTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/transfer', [
            'sin_to' => 7700000,
            'amount' => 100000000,
            'comment' => 'Testtransfer'
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::UNPROCESSABLE_ENTITY);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([]);
        $I->seeResponseMatchesJsonType([
            'data' => 'null',
            'message' => 'string',
            'status' => 'boolean',
        ]);

        $I->canSeeResponseContainsJson([
            'status' => false,
        ]);
    }

    public function createTrasferWithoutAuthFailTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendPOST(self::$route . '/transfer', [
            'created_at' => '2019-06-04T06:21:20.456Z',
            'sin_from' => 80,
            'sin_to' => 1,
            'amount' => 100,
            'comment' => 'string',
            'recurrent_payment_id' => 0
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::UNPROCESSABLE_ENTITY);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'data' => 'null',
            'message' => 'string',
            'status' => 'boolean',
        ]);

        $I->canSeeResponseContainsJson([
            'status' => false,
        ]);
    }
}
