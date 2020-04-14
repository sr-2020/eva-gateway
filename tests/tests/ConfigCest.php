<?php

class ConfigCest
{
    static protected $route = '/config';
    static protected $testSet;

    public function setTest(ApiTester $I)
    {
        $rand = rand(0, 1000000);
        self::$testSet = [
            'config' => 'set',
            'value' => 'value' . $rand,
        ];

        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendPOST(self::$route . '/test', self::$testSet);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::CREATED);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'config' => 'string',
            'value' => 'string'
        ]);
    }

    public function getTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendGET(self::$route . '/test');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'config' => 'string',
            'value' => 'string'
        ]);

        $I->canSeeResponseContainsJson([
            'value' => self::$testSet['value'],
        ]);
    }

    public function getNotFound(ApiTester $I)
    {
        $I->sendGET(self::$route . '/test404');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::NOT_FOUND);
    }
}
