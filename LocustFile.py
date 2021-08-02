
from locust import HttpUser, TaskSet, task, between


class UserBehavior(TaskSet):

    @task(1)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def createGiftCodes(self):

        data = {
            "jsonStr": "{'createUserId':'1000','giftDescribe':'礼品码1','giftList':[{'name':'士兵','num':10},{'name':'金币','num':10}],'giftCodeType':'A','validityStr':'2021-08-30','giftPullUser':'100001'}"

        }
        response = self.client.post('/createGiftCodes', data = data,name="createGiftCodes")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')
    @task(2)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def getCiftCodes(self):

        data = {
            "giftCode": "ZHTDN52L"

        }
        response = self.client.post('/getCiftCodes', data = data,name="getCiftCodes")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')
    @task(3)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def activateCode(self):

        data = {
            "giftCode": "ZHTDN52L",
            "userId":"100017"

        }
        response = self.client.post('/activateCode', data = data,name="activateCode")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')



class TestLocust(HttpUser):

    tasks = [UserBehavior]
    wait_time = between(2, 5)
    host = "http://127.0.0.1:8000"
    #task_set = UserBehavior
    #host = "http://127.0.0.1/:8000"  # 被测服务器地址
    #min_wait = 5000
# 最小等待时间，即至少等待多少秒后Locust选择执行一个任务。

    #max_wait = 9000
# 最大等待时间，即至多等待多少秒后Locust选择执行一个任务。