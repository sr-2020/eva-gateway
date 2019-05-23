<?php

namespace Helper;

/**
 * Description of ManyToManyTrait
 *
 * @author ognestraz
 */
trait ManyToManyTrait
{
    public function _before(\ApiTester $I)
    {
        if (static::$before) {
            return ;
        }
        $name = 'Test';
        $I->haveHttpHeader('Content-Type', 'application/json');
        $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
        $I->sendPOST(self::$route, ['name' => $name]);
        $jsonResponse = json_decode($I->grabResponse());
        self::$createdId = $jsonResponse->id;
        
        for ($i = 0; $i < self::$countItems; $i++) {
            $name = 'Sub Test' . $i;
            $I->haveHttpHeader('Content-Type', 'application/json');
            $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
            $I->sendPOST(self::$subroute, ['name' => $name]);
            $jsonResponse = json_decode($I->grabResponse(), true);
            self::$createdSubs[] = $jsonResponse;            
        }
        static::$before = true;
    }
    
    public function indexTest(\ApiTester $I)
    {
        $I->sendGET(self::$route . '/' . self::$createdId . self::$subroute);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        $I->seeResponseContainsJson([]);
    }

    public function createTest(\ApiTester $I)
    {
        foreach (self::$createdSubs as $sub) {
            $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
            $I->sendPOST(self::$route . '/' . self::$createdId . self::$subroute . '/' . $sub['id']);
        }

        $I->sendGET(self::$route . '/' . self::$createdId . self::$subroute);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();
        
        $I->seeResponseContainsJson(self::$createdSubs);
    }

    public function deleteTest(\ApiTester $I)
    {
        foreach (self::$createdSubs as $sub) {
            $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
            $I->sendDELETE(self::$route . '/' . self::$createdId . self::$subroute . '/' . $sub['id']);
        }

        $I->sendGET(self::$route . '/' . self::$createdId . self::$subroute);
        $I->seeResponseCodeIs(\Codeception\Util\HttpCode::OK);
        $I->seeResponseIsJson();

        $I->seeResponseContainsJson([]);

        $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
        $I->sendDELETE(self::$route . '/' . self::$createdId);
        foreach (self::$createdSubs as $sub) {
            $I->haveHttpHeader('Authorization', 'Token ' . $I->getAdminToken());
            $I->sendDELETE(self::$subroute . '/' . $sub['id']);
        }
    }
}
