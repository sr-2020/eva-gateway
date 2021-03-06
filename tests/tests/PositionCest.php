<?php

class PositionCest
{
    static protected $route = '/position';

    static protected $data;

    public function positionsListTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/positions');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
    }

    public function createPositionTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST(self::$route . '/positions', [
            'beacons' => [
                [
                    'ssid' => 'beacon1',
                    'bssid' => 'b0:0a:95:9d:00:0a',
                    'level' => -50
                ]
            ]
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::CREATED);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'user_id' => 'integer',
            'created_at' => 'string',
        ]);
    }

    public function createPositionFailTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer test');
        $I->sendPOST(self::$route . '/positions', [
            'beacons' => [
                [
                    'ssid' => 'beacon1',
                    'bssid' => 'b0:0a:95:9d:00:0a',
                    'level' => -50
                ]
            ]
        ]);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::UNAUTHORIZED);
    }
}
