import { Injectable } from '@nestjs/common';
import { FinalTorrent } from './interfaces/torrent.interface';
import * as torrentApi from 'torrent-search-api';

torrentApi.enableProvider('1337x');

@Injectable()
export class AppService {

  async getLongMagnet(desc: string): Promise<string> {
    //@ts-ignore
    return await torrentApi.getMagnet({provider: "1337x", desc: desc });
  }

  async getTorrents(torrent: string): Promise<FinalTorrent[]> {
    try {
      const torrents: FinalTorrent[] = await torrentApi.search(torrent);
      const finalTorrents: FinalTorrent[] = [];

      for (let i = 0; i < 10; i++) {
        if (torrents[i] != null) {
          finalTorrents.push(torrents[i]);
        }
      }

      for (let i = 0; i < finalTorrents.length; i++) {
        const longMagnet: string = await torrentApi.getMagnet(finalTorrents[i]);
        finalTorrents[i].magnet = longMagnet;
        finalTorrents[i].number = i + 1;
      }

      return finalTorrents;
    } catch (error: any) {
      console.log(error);
      return [];
    }
  }
}
