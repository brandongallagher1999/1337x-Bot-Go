import { Controller, Get, Param } from '@nestjs/common';
import { AppService } from './app.service';
import { FinalTorrent, MagnetResponse } from './interfaces/torrent.interface';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get(':torrent')
  async getTorrents(@Param() params): Promise<FinalTorrent[]> {
    return await this.appService.getTorrents(params.torrent);
  }

  @Get("/longMagnet/:longMagnet")
  async getLongMagnet(@Param() params): Promise<MagnetResponse> {
    return { magnet: await this.appService.getLongMagnet(params.longMagnet) };
  }
}
