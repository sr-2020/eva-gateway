<?php

class AuthCest
{
    static protected $createdId = 0;
    static protected $route = '/auth';

    static private $testCreds = [];
    static private $createdApiKey = '';

    public function profileTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendGET(self::$route . '/profile');
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
        $I->sendPOST(self::$route . '/register', self::$testCreds);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::CREATED);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'api_key' => 'string'
        ]);

        $jsonResponse = json_decode($I->grabResponse());
        self::$createdApiKey = $jsonResponse->api_key;
        self::$createdId = $jsonResponse->id;
    }

    public function profileAfterRegisterTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . self::$createdApiKey);
        $I->sendGET(self::$route . '/profile');
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
        $I->sendPOST(self::$route . '/login', self::$testCreds);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'api_key' => 'string'
        ]);

        $jsonResponse = json_decode($I->grabResponse());
        self::$createdApiKey = $jsonResponse->api_key;
    }

    public function loginWithFirebaseTokenTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        self::$testCreds['firebase_token'] = 'testtoken' . rand(0, 10000000);
        $I->sendPOST(self::$route . '/login', self::$testCreds);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseMatchesJsonType([
            'id' => 'integer',
            'api_key' => 'string'
        ]);

        $jsonResponse = json_decode($I->grabResponse());
        self::$createdApiKey = $jsonResponse->api_key;
    }

    public function pushNotificationAfterLoginWithFirebaseTokenTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . $I->getToken());
        $I->sendPOST('/push/send_notification/' . self::$createdId, [
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
        $I->seeResponseContainsJson([
            'token_used' => self::$testCreds['firebase_token'],
            'failure' => 1,
            'success' => 0,
        ]);
    }

    public function profileAfterLoginTest(ApiTester $I)
    {
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Bearer ' . self::$createdApiKey);
        $I->sendGET(self::$route . '/profile');
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
        $I->sendPUT(self::$route . '/profile', $editData);
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
        $I->sendGET(self::$route . '/profile');
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::UNAUTHORIZED);
    }
}
