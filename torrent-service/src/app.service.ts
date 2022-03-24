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

      return finalTorrents;
    } catch (error: any) {
      console.log(error);
      return [];
    }
  }
}
