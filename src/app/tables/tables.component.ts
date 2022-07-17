import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

declare interface TableData {
    headerRow: string[];
    dataRows: string[][];
}

@Component({
  selector: 'app-tables',
  templateUrl: './tables.component.html',
  styleUrls: ['./tables.component.css']
})
export class TablesComponent {
    public tableData1: TableData;
    public tableData2: TableData;
    pokemon:any
    random:number

    constructor(private http: HttpClient) { }
    listaPokemones:any = []
  

    loadPokemones(){
    this.http
    .get("http://localhost:8000/pokemones")
    .subscribe((listaPokemones: any) => {
        this.listaPokemones = listaPokemones
    //.subscribe((response) => {
        //alert(JSON.stringify(response));
    })
    }
    loadPokemonRandom(){
        this.random = this.generateRandom(1,151)
        //alert(this.random)
        this.http
        .get("http://localhost:8000/pokemones/"+this.random)
        .subscribe((listaPokemones: any) => {
          this.listaPokemones = listaPokemones
        //.subscribe((response) => {
          //alert(JSON.stringify(response));
        })
    }
    generateRandom(min: number, max:number){
       
        return Math.floor(Math.random() * (max - min + 1)) + min; 

    }
  


}
