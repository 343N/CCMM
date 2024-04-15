import { os } from "@tauri-apps/api";
import { BaseDirectory, createDir, writeFile } from "@tauri-apps/api/fs";
import { AppStore } from "./AppStore";
import { getClient } from "@tauri-apps/api/http";

let DEFAULT;



class ModFileManager {

  static LISTENER = {
    DOWNLOAD_STARTED: 0,
    DOWNLOAD_FINISHED: 1,
    DOWNLOAD_ERROR: 2,
    MOD_INSTALLED: 3,
    MOD_DISABLED: 4,
    MOD_HAS_UPDATE: 5,
  }

  constructor(modDir="downloads"){
    this.downloadDir = modDir
    this.makeDownloadDir()
    this.listeners = []
    for (let key of LISTENER) 
      this.listeners.push([])
  }


  async downloadMod(modinfo){
    let get = await getClient()

    this.fireEvent(DOWNLOAD_STARTED, modinfo)
    let resp = await get.request({
      method: 'GET',
      url: modinfo.modfile.download
    })
    if (resp.ok)
      this.fireEvent(DOWNLOAD_FINISHED, modinfo)
    else this.fireEvent(DOWNLOAD_ERROR, modinfo, resp.status)


  }

  fireEvent(type, ...args){
    for (let fn of this.listeners[type]){
      fn(...args)
    }
  }

  addListener(type, fn){
    this.listeners[type].push(fn)
  }
  
  
  async installMod(modinfo, filename){
  
  }

  async isModInstalled(modinfo){

  }

  async checkDownloadedForUpdates(){

  }

  async checkModHasUpdates(){

  }

  async checkForOfflineMods(){

  }

  async getModFileHash(){

  }

  async getModuleFolderHash(){

  }
  
  async makeDownloadDir(dir){
    await createDir(dir)
  }
  
  saveDownloadData(){
  
  }
}


export default ModFileManager;
