<template>
  <div class="click-to-copy">

    <v-text-field
      v-if="input"
      :id="itemId + uniqueNumber"
      :readonly="isReadOnly"
      :value="inputValue"
      :class="inputMargin"
      :hint="copied ? 'copied' : ''"
      append-icon="content_copy"
      filled
      class="click-to-copy__input"
      @click:append.stop="handleCopy(itemId + uniqueNumber)"
    >
    </v-text-field>
    <div
      v-resize="onResize"
      v-else
      class="click-to-copy__no-input"
      @click.self="clicked"
    >
      <span v-if="windowWidth < 1264 || crop">
        <v-tooltip bottom color="black">
          <template v-slot:activator="{ on }">
            <span
              class="click-to-copy__crop"
              @click="clicked"
              v-on="on"
            >
              {{ inputValue | crop }}
            </span>
          </template>
          <span @click="clicked">{{ inputValue }}</span>
        </v-tooltip>
      </span>
      <div v-else class="click-to-copy__no-input--large" @click="clicked">
        {{ inputValue }}
      </div>
      <v-tooltip :color="copied ? 'accent' : 'black'" class="click-to-copy__tooltip" bottom>
        <template v-slot:activator="{ on }">
          <i
            id="copy-icon"
            class="material-icons"
            v-on="on"
            @mousedown.stop="handleCopy(itemId + uniqueNumber)"
            @mouseleave="handleTooltip"
          >
            content_copy
          </i>
        </template>
        <span :class="{ 'text--black': copied }">
          {{ copied ? 'Copied!' : 'Copy ID' }}
        </span>
      </v-tooltip>
    </div>

    <!-- USAGE:
    <Click-To-Copy
      :input-value="props.item.tx_id"
      :is-read-only="true"
      :input="false" // determines if the value will be in an input element or div
      :crop="true" // determines if value should be cropped, only on non-input items
      input-margin="ma-1" // String which determines the margin class which should be applied to the input
      item-id="txid"
      @ctcClicked="someFunction"
    ></Click-To-Copy> -->
  </div>
</template>

<script>
export default {
  name: 'ClickToCopy',
  props: {
    itemId: {
      required: true,
      type: String,
      default: (Math.floor(Date.now() * Math.random())).toString()
    },
    inputValue: {
      type: [String, Number],
      default: ''
    },
    isReadOnly: {
      type: Boolean,
      default: false
    },
    input: {
      type: Boolean,
      default: true
    },
    inputMargin: {
      type: String,
      default: 'ma-0'
    },
    crop: {
      type: Boolean,
      default: false
    }
  },
  data: () => ({
    copied: false,
    copySucceeded: false,
    copyFailed: false,
    uniqueNumber: (Math.floor(Date.now() * Math.random())).toString(),
    windowWidth: Number
  }),
  mounted () {
    this.onResize()
  },
  methods: {
    handleCopy (id) {
      if (document.execCommand) {
        const copyText = document.getElementById(id)
        if (this.input) {
          copyText.select()
          document.execCommand('copy')
        } else {
          navigator.clipboard.writeText(this.inputValue)
        }
        this.copied = true
      } else {
        this.onError()
      }
    },
    clicked () {
      this.$emit('ctcClicked')
    },
    handleTooltip () {
      setTimeout(() => {
        this.copied = false
      }, 300)
    },
    onResize () {
      this.windowWidth = window.innerWidth
    }
  }
}
</script>

<style>
.click-to-copy {
  position: relative;
  display: flex;
  flex-grow: 1;
}
.click-to-copy__input {
  position: relative;
}
/* .click-to-copy__input .input-group__input input {
  padding-right: 30px !important;
} */
.click-to-copy__no-input {
  display: flex;
  justify-content: flex-start;
  width: 100%;
}
.click-to-copy__no-input--large {
  min-width: 520px;
}
.click-to-copy__crop {
  white-space: nowrap;
}
.click-to-copy__tooltip {
  position: absolute;
  top: 15px;
  font-size: 16px;
}
.click-to-copy__alert-wrapper {
  position: absolute;
  right: 0;
  top: 0;
  margin-top: 58px;
}
.click-to-copy__alert-wrapper--no-input {
  margin-top: 18px;
}
.click-to-copy .copy--alert {
  position: absolute;
  opacity: 0;
  font-size: 12px;
  margin-bottom: 10px;
  top: 0px;
  right: 0;
  white-space: nowrap;
}
.click-to-copy .copy--success {
  color: $success;
}
.click-to-copy .copy--fail {
  color: $error;
}
.click-to-copy #copy-icon {
  padding: 4px 0 0 10px;
  font-size: 20px;
  line-height: 14px;
}
.click-to-copy #copy-icon:hover {
  cursor: pointer;
}
.click-to-copy .copyAlert--success {
  animation: copyAlert 2s ease-in-out;
}
.click-to-copy .copyAlert--fail {
  animation: copyAlert 5s ease-in-out;
}
.click-to-copy .v-messages {
  text-align: right;
}
@-moz-keyframes copyAlert {
  0% {
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
@-webkit-keyframes copyAlert {
  0% {
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
@-o-keyframes copyAlert {
  0% {
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
@keyframes copyAlert {
  0% {
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}

</style>