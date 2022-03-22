import { Test, TestingModule } from '@nestjs/testing';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { FinalTorrent } from './interfaces/torrent.interface';

describe('AppController', () => {
  let appController: AppController;

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService],
    }).compile();

    appController = app.get<AppController>(AppController);
  });

  describe('root', () => {
    it('should return "Hello World!"', async () => {
      console.time();
      const res: FinalTorrent[] = await appController.getTorrents('Inception'); //first object in response
      console.log(res);
      expect(Object.keys(res[0])).toStrictEqual([
        'title',
        'time',
        'seeds',
        'peers',
        'size',
        'desc',
        'provider',
        'magnet',
        'number',
      ]); //verifying schema
      console.timeEnd();
    });
  });
});
