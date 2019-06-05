<?php

class UserCest
{
    static protected $createdId = 0;
    static protected $route = '/users';

    static private $testCreds = [];
    static private $createdApiKey = '';

    public function indexTest(ApiTester $I)
    {
        $I->sendGET(self::$route);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->canSeeResponseContainsJson([]);
    }

    public function profileTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET('/auth/profile');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'name' => 'string',
            'status' => 'string',
            'created_at' => 'string',
            'updated_at' => 'string',
            'admin' => 'boolean',
        ]);
    }

    public function registerTest(ApiTester $I)
    {
        $rand = rand(0, 1000000);
        self::$testCreds = [
            'email' => 'email' . $rand . '@mail.com',
            'name' => 'name' . $rand,
            'password' => 'password' . $rand,
        ];

        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendPOST('/auth/register', self::$testCreds);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::CREATED);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'api_key' => 'string'
        ]);

        $jsonResponse = json_decode($I->grabResponse());
        self::$createdApiKey = $jsonResponse->api_key;
    }

    public function profileAfterRegisterTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . self::$createdApiKey);
        $I->sendGET('/auth/profile');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'name' => 'string',
            'status' => 'string',
            'created_at' => 'string',
            'updated_at' => 'string',
            'admin' => 'boolean',
        ]);

        $I->canSeeResponseContainsJson([
            'name' => self::$testCreds['name'],
        ]);
    }

    public function loginTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->sendPOST('/auth/login', self::$testCreds);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'api_key' => 'string'
        ]);

        $jsonResponse = json_decode($I->grabResponse());
        self::$createdApiKey = $jsonResponse->api_key;
    }

    public function profileAfterLoginTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . self::$createdApiKey);
        $I->sendGET('/auth/profile');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'name' => 'string',
            'status' => 'string',
            'created_at' => 'string',
            'updated_at' => 'string',
            'admin' => 'boolean',
        ]);

        $I->canSeeResponseContainsJson([
            'name' => self::$testCreds['name'],
        ]);
    }

    public function editProfileTest(ApiTester $I)
    {
        $editData = [
            'status' => 'platinum',
        ];

        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . self::$createdApiKey);
        $I->sendPUT('/auth/profile', $editData);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'name' => 'string',
            'status' => 'string',
            'created_at' => 'string',
            'updated_at' => 'string',
            'admin' => 'boolean',
        ]);

        $I->canSeeResponseContainsJson([
            'name' => self::$testCreds['name'],
            'status' => $editData['status'],
        ]);
    }

    public function profileFailTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer test');
        $I->sendGET('/auth/profile');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::UNAUTHORIZED);
    }
}
