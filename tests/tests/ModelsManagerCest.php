<?php

class ModelsManagerCest
{
    static protected $route = '/models-manager';

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
            'url' => 'string',
            'headers' => 'array',
        ]);
    }

    public function characterModelNotFoundTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/character/model');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::NOT_FOUND);
        $I->seeResponseContains("Character model with id = 1 not found");
    }
}
