<template>
	<div class="home-container layout-pd">
		<el-row :gutter="15" class="home-card-one mb15">
      <el-col :span="24" class="home-media">
        <div class="home-card-item">
          <div class="home-card-item-main-title">
            {{ "消息发送服务已就绪，随时待命！" }}
          </div>
        </div>
      </el-col>
		</el-row>
		<el-row :gutter="15" class="home-card-two">
			<el-col :xs="24" :sm="8" class="home-media mb15">
				<div class="home-card-item">
          <div class="home-card-item-title">
            {{ "消息平台" }}
          </div>
          <el-row>
            <el-col
                :xs="12" :sm="24"
                v-for="(value,key) in state.providers" :key="key">
              <div class="mt20 ml10 mr10 mb20">
                <div class="gradient-border" :style="{ borderColor: colorMap.get(key), boxShadow: `0 4px 12px ${colorMapLite.get(key)}` }">
                  <div
                      style="font-size: 20px"
                      class="mt5 ml5"
                  >{{$t('message.category.'+key)}}</div>
                  <el-descriptions
                      class="mt15 ml5"
                      :column="2"
                      direction="vertical">
                    <el-descriptions-item
                        v-for="item in value"
                        :key="item.provider"
                        :label="item.provider"
                        class="home-card-item-small-title">
                      累计发送 {{ item.count }} 条
                    </el-descriptions-item>
                  </el-descriptions>
                </div>
              </div>
            </el-col>
          </el-row>
				</div>
			</el-col>
			<el-col :xs="24" :sm="16" class="mb15">
				<div class="home-card-item">
					<div style="height: 100%" ref="homeLineRef"></div>
				</div>
			</el-col>
		</el-row>
	</div>
</template>

<script setup lang="ts" name="home">
import { reactive, onMounted, ref, watch, nextTick, onActivated, markRaw } from 'vue';
import * as echarts from 'echarts';
import { storeToRefs } from 'pinia';
import { useThemeConfig } from '/@/stores/themeConfig';
import { useTagsViewRoutes } from '/@/stores/tagsViewRoutes';
import {
  getHomeOne,
  getHomeThree,
} from "/@/views/home/mock";
import {getProviderStatApi, getDailyStatApi} from "/@/api/history";

// 定义变量内容
const homeLineRef = ref();
const storesTagsViewRoutes = useTagsViewRoutes();
const storesThemeConfig = useThemeConfig();
const { themeConfig } = storeToRefs(storesThemeConfig);
const { isTagsViewCurrenFull } = storeToRefs(storesTagsViewRoutes);
const state = reactive({
	global: {
		homeChartOne: null,
		homeChartTwo: null,
		homeCharThree: null,
		dispose: [null, '', undefined],
	} as any,
	homeOne: getHomeOne(),
	homeThree: getHomeThree(),
	myCharts: [] as EmptyArrayType,
	charts: {
		theme: '',
		bgColor: '',
		color: '#303133',
	},
  providers: {}
});

const mailColor = 'rgba(108,80,243,0.3)';
const smsColor = 'rgba(64,186,209,0.3)';
const imColor = 'rgba(185,91,216,0.3)';
const colors = [mailColor, smsColor, imColor];
const colorMapLite = new Map([
  ["mail", mailColor],
  ["sms", smsColor],
  ["im", imColor],
])
const colorMap = new Map([
  ["mail", 'rgba(108,80,243,0.6)'],
  ["sms", 'rgba(64,186,209,0.6)'],
  ["im", 'rgba(185,91,216,0.6)'],
])

let rawData : dailyStat[] = [];

interface dailyStat {
  date: string;
  mail: number;
  sms: number;
  im: number;
  total: number;
}

const getVerticalBarOption = (color: string, bgColor: string) => {
  const data = rawData.reverse();
  const yAxis = data.map(item => item.date);
  const totalData = data.map(item => item.total);

  return {
    backgroundColor: bgColor,
    title: {
      text: '发送历史',
      x: 'left',
      textStyle: {fontSize: '15', color: color},
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    grid: {
      right: 80 // 预留 dataZoom 空间
    },
    xAxis: {
      type: 'value',
      boundaryGap: [0, 0.01]
    },
    yAxis: {
      type: 'category',
      data: yAxis,
      axisLabel: {
        show: true
      },
      splitLine: {
        show: false
      }
    },
    color: colors,
    dataZoom: [
      {
        type: 'slider',
        start: 80,
        end: 100, // 初始显示，百分比
        orient: 'vertical',
        right: 10,
        top: 30,
        bottom: 30,
        height: 200,
        realtime: true,
        zoomLock: false,
        filterMode: 'filter'
      }
    ],
    series: [
      {
        name: 'mail',
        type: 'bar',
        data: data.map(item => item.mail),
        stack: '总量',
        animationEasing: "quarticOut",
        animationDuration: 500,
        itemStyle: {
          //柱状图圆角
          borderRadius: [5, 5, 5, 5],
        },
      },
      {
        name: 'sms',
        type: 'bar',
        data: data.map(item => item.sms),
        stack: '总量',
        animationEasing: "quarticOut",
        animationDuration: 500,
        // animationDelay: 100,
        itemStyle: {
          //柱状图圆角
          borderRadius: [5, 5, 5, 5],
        },
      },
      {
        name: 'im',
        type: 'bar',
        data: data.map(item => item.im),
        animationDuration: 500,
        // animationDelay: 200,
        stack: '总量',
        animationEasing: "quarticOut",
        itemStyle: {
          //柱状图圆角
          borderRadius: [5, 5, 5, 5],
        },
        label: {
          show: true,
          position: 'right',
          valueAnimation: true,
          formatter: ({ dataIndex }) => totalData[dataIndex], // 显示总值
        },
      }
    ]
  }
}

// 每日记录
const initBarChart = () => {
	if (!state.global.dispose.some((b: any) => b === state.global.homeChartOne)) state.global.homeChartOne.dispose();
	state.global.homeChartOne = markRaw(echarts.init(homeLineRef.value, state.charts.theme));
	const option = getVerticalBarOption(state.charts.color, state.charts.bgColor, rawData);
	state.global.homeChartOne.setOption(option);
	state.myCharts.push(state.global.homeChartOne);
  homeLineRef.value.addEventListener('wheel', handleWheel)
  // state.global.homeChartOne.getZr().on('wheel', handleWheel)
};

let loading = false;

// 处理滚轮事件
function handleWheel(event) {
  if (loading) return;
  const option = state.global.homeChartOne.getOption()
  const currentZoom = option.dataZoom[0]
  let { start, end } = currentZoom

  const windowSize = end - start

  const delta = event.delta || (event.wheelDelta ? -event.wheelDelta : 0)

  if (delta < 0) {
    // 向下滚动
    start += 1
  } else {
    // 向上滚动
    start -= 1
  }

  // 限制边界
  start = Math.max(0, Math.min(start, 100 - windowSize))
  end = start + windowSize

  // 更新 dataZoom 范围
  state.global.homeChartOne.setOption({
    dataZoom: [{ start, end }]
  })
}

// 批量设置 echarts resize
const initEchartsResizeFun = () => {
	nextTick(() => {
		for (let i = 0; i < state.myCharts.length; i++) {
			setTimeout(() => {
				state.myCharts[i].resize();
			}, i * 1000);
		}
	});
};
// 批量设置 echarts resize
const initEchartsResize = () => {
	window.addEventListener('resize', initEchartsResizeFun);
};

const getCurrentDate = () => {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0');
  const day = String(now.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}

// 页面加载时
onMounted(() => {
  getProviderStatApi({}).then(res => {
    if (res && res.code === 200 && res.providers) {
      state.providers = res.providers;
    }
  })
	initEchartsResize();
});
// 由于页面缓存原因，keep-alive
onActivated(() => {
	initEchartsResizeFun();
});
// 监听 pinia 中的 tagsview 开启全屏变化，重新 resize 图表，防止不出现/大小不变等
watch(
	() => isTagsViewCurrenFull.value,
	() => {
		initEchartsResizeFun();
	}
);
// 监听 pinia 中是否开启深色主题
watch(
	() => themeConfig.value.isIsDark,
	(isIsDark) => {
		nextTick(() => {
			state.charts.theme = isIsDark ? 'dark' : '';
			state.charts.bgColor = isIsDark ? 'transparent' : '';
			state.charts.color = isIsDark ? '#dadada' : '#303133';
      getDailyStatApi({
        date: getCurrentDate(),
        batch: 100,
        min: 100,
      }).then(res => {
        rawData = res.stats;
        initBarChart();
      })

		});
	},
	{
		deep: true,
		immediate: true,
	}
);
</script>

<style scoped lang="scss">
$homeNavLengh: 8;
.home-container {
	overflow: hidden;
  display: flex;
  flex-direction: column;
  flex: 1;
	.home-card-one,
	.home-card-two,
	.home-card-three {
		.home-card-item {
			width: 100%;
			height: 80px;
			border-radius: 4px;
			transition: all ease 0.3s;
			padding: 20px;
			overflow: hidden;
			background: var(--el-color-white);
			color: var(--el-text-color-primary);
			border: 1px solid var(--next-border-color-light);
			&:hover {
				box-shadow: 0 2px 12px var(--next-color-dark-hover);
				transition: all ease 0.3s;
			}
			&-icon {
				width: 70px;
				height: 70px;
				border-radius: 100%;
				flex-shrink: 1;
				i {
					color: var(--el-text-color-placeholder);
				}
			}
			&-title {
				font-size: 15px;
				font-weight: bold;
				height: 30px;
			}
      &-main-title {
        font-size: 27px;
        font-weight: bold;
        color: var(--el-color-primary);
        height: 30px;
      }
      &-small-title {
        color: var(--el-color-primary);
      }
		}
	}
	.home-card-two,
	.home-card-three {
    flex: 1;
		.home-card-item {
			height: 100%;
			width: 100%;
			overflow: hidden;
			.home-monitor {
				height: 100%;
				.flex-warp-item {
					width: 25%;
					height: 111px;
					display: flex;
					.flex-warp-item-box {
						margin: auto;
						text-align: center;
						color: var(--el-text-color-primary);
						display: flex;
						border-radius: 5px;
						background: var(--next-bg-color);
						cursor: pointer;
						transition: all 0.3s ease;
						&:hover {
							background: var(--el-color-primary-light-9);
							transition: all 0.3s ease;
						}
					}
					@for $i from 0 through $homeNavLengh {
						.home-animation#{$i} {
							opacity: 0;
							animation-name: error-num;
							animation-duration: 0.5s;
							animation-fill-mode: forwards;
							animation-delay: calc($i/10) + s;
						}
					}
				}
			}
		}
	}
}
.gradient-border {
  padding: 10px;
  border: 2px solid;
  border-radius: 12px;
}
</style>
