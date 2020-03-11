<?php

class ModelEngineCest
{
    static protected $route = '/model-engine';

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

}
