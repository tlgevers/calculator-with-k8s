<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <!-- <q-btn -->
        <!--   flat -->
        <!--   dense -->
        <!--   round -->
        <!--   icon="menu" -->
        <!--   aria-label="Menu" -->
        <!--   @click="leftDrawerOpen = !leftDrawerOpen" -->
        <!-- /> -->

        <q-toolbar-title>
          Quasar App
        </q-toolbar-title>

        <!-- <div>Quasar v{{ $q.version }}</div> -->
      </q-toolbar>
    </q-header>

    <!-- <q-drawer -->
    <!--   v-model="leftDrawerOpen" -->
    <!--   show-if-above -->
    <!--   bordered -->
    <!--   content-class="bg-grey-1" -->
    <!-- > -->
    <!--   <q-list> -->
    <!--     <q-item-label -->
    <!--       header -->
    <!--       class="text-grey-8" -->
    <!--     > -->
    <!--       Essential Links -->
    <!--     </q-item-label> -->
    <!--     <EssentialLink -->
    <!--       v-for="link in essentialLinks" -->
    <!--       :key="link.title" -->
    <!--       v-bind="link" -->
    <!--     /> -->
    <!--   </q-list> -->
    <!-- </q-drawer> -->

    <q-page-container>
      <div class="column" style="width: 320px;">
        <q-field filled label="Filled" stack-label>
          <template v-slot:control>
            <div class="self-center full-width no-outline" tabindex="0">{{screenValue}}</div>
          </template>
        </q-field>
      </div>
      <div class="q-pa-md row">
        <div class="row" style="height: 260px; max-width: 185px;">
          <Key
            @getKey="getKey"
            v-for="key in calcKeys"
            :key="key.title"
            v-bind="key"
          />
        </div>
        <div class="row" style="height: 260px; max-width: 65px;">
          <Key
            v-for="operate in operators"
            :key="operate.title"
            v-bind="operate"
          />
        </div>
        <div class="row" style="height: 260px; max-width: 65px;">
          <Key
            v-for="comp in compute"
            :key="comp.title"
            v-bind="comp"
          />
        </div>
      </div>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
// import EssentialLink from 'components/EssentialLink.vue'
import Key from 'components/button.vue'

const Keys = [
  {
    title: '1',
    classCustom: 'col-4'
  },
  {
    title: '2',
    classCustom: 'col-4'
  },
  {
    title: '3',
    classCustom: 'col-4'
  },
  {
    title: '4',
    classCustom: 'col-4'
  },
  {
    title: '5',
    classCustom: 'col-4'
  },
  {
    title: '6',
    classCustom: 'col-4'
  },
  {
    title: '7',
    classCustom: 'col-4'
  },
  {
    title: '8',
    classCustom: 'col-4'
  },
  {
    title: '9',
    classCustom: 'col-4'
  },
  {
    title: '0',
    classCustom: 'col-4'
  },
  {
    title: '.',
    classCustom: 'col-4'
  },
  {
    title: '%',
    classCustom: 'col-4'
  }
]

const Operators = [
  {
    title: '/',
    classCustom: '',
    bgColor: 'grey'
  },
  {
    title: 'X',
    classCustom: '',
    bgColor: 'grey'
  },
  {
    title: '-',
    classCustom: '',
    bgColor: 'grey'
  },
  {
    title: '+',
    classCustom: '',
    bgColor: 'grey'
  }
]

const Compute = [
  {
    title: '=',
    classCustom: '',
    bgColor: 'orange'
  }
]

export default {
  name: 'MainLayout',
  components: { Key },
  data () {
    return {
      leftDrawerOpen: false,
      calcKeys: Keys,
      operators: Operators,
      compute: Compute,
      screenValue: ''
    }
  },
  methods: {
    getKey (key) {
      this.screenValue = this.screenValue + key
      console.log('key pressed', key)
    }
  }
}
</script>
