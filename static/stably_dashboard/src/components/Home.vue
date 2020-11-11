<template>
  <v-container>
    <v-row class="text-center">
      <v-col
        class="mb-5"
        cols="12"
      >
        <v-card light
          class="mx-auto my-12"
          max-width="450"
        >
        <v-img
          height="200"
          src="../assets/header.jpg"
        ></v-img>
          <v-card-title>Find Prime Number</v-card-title>
          <v-card-text>
          <v-container>
            <v-row>
              <v-col
                cols="8"
              >
                <v-text-field
                    label="Input Number"
                    type="number"
                    v-model="number"
                  ></v-text-field>
              </v-col>
              <v-col>
                <v-btn
                  depressed
                  color="primary"
                  class="mt-3"
                  @click.native="findnumber"
                >
                  Find
                </v-btn>
              </v-col>
            </v-row>
            <v-card left height="70">
              <v-card-text>
                <p class="text-sm-left">Result: {{result}}</p>
              </v-card-text>
            </v-card>
          </v-container>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import {findPrime} from '../api/prime'
  export default {
    name: 'Home',
    data: () => ({
      number: 0,
      result: ''
    }),
    mounted(){},
    methods:{
      findnumber(){
        findPrime(this.number).then(res=>{
          if(res.data===0){
            this.result='Not found the highest prime number lower than the input number!';
          }else{
            this.result='The highest prime number lower than the input number: ' + res.data;
          }
        }).catch(err=>{
          this.result = JSON.stringify(err)
        })
      }
    }
  }
</script>
