<?php

class BillingCest
{
    static protected $route = '/account_info';

    static protected $data;

    public function readTest(ApiTester $I)
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
}
