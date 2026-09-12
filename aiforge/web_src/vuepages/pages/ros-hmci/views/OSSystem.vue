<template>
  <div class="flex-col page">
    <AppBanner></AppBanner>

    <div class="flex-col justify-start items-center relative section_6">
      <div class="flex-col space-y-72">
        <div class="flex-col justify-start self-center relative group_3">
          <span class="text_16 pos_99">关于人机协同智能操作系统</span>
          <!-- <span class="font_3 text_14 text_15 pos_4">Brief introduction</span> -->
        </div>
        <div class="flex-row space-x-92">
          <div ref="leftElem" class="flex-col self-start group_5 space-y-54 slide-in-left" :class="{ 'show': showLeft }">
            <div class="flex-col">
              <div class="flex-row space-x-16">
                <img class="image_12" src="/img/ros-hmci/mbz606.png" />
                <span class="font_41 text_17">系统概述</span>
              </div>
              <div class="flex-col items-center group_6 space-y-2">
                <span class="font_5">人机协同智能操作系统是一组软件库和工具，可帮助您构建人工智能应用程序。
                  从驱动程序到最先进的算法，海量的数据集，模型，以及强大的开发人员工具...
                </span>
              </div>
              <div class="flex-row group_7 space-x-8">
                <div class="flex-col justify-start text-wrapper_2">
                  <span class="font_6 text_18">人机协同智能操作系统...</span>
                </div>
                <div class="flex-col justify-start text-wrapper_2">
                  <span class="font_6 text_18">复杂环境与作业规划器</span>
                </div>
              </div>
            </div>
            <div class="flex-row items-center self-start section_8">
              <a :href="links.os_info" class="text_19">更多详情</a>
              <img class="shrink-0 image_14" src="/img/ros-hmci/mbz607.png" />
            </div>
          </div>
          <img ref="rightElem" class="image_27 slide-in-right" :class="{ 'show': showRight }"
            src="/img/ros-hmci/group4654.png" />
        </div>
      </div>
    </div>
    <div class="flex-col justify-start relative">
      <div class="flex-col justify-start section_9 pos_9">
        <div class="flex-col justify-start items-center relative section_30">
          <img class="image_1" src="/img/ros-hmci/6da4828190161924d3e2c084e55fa79e.png" />
          <img class="image_13 pos_1" src="/img/ros-hmci/95ecb911f420e03fb3ade00d4bc53614.png" />
        </div>
      </div>
      <div class="flex-col relative">
        <div class="flex-col justify-start items-center group_1">
          <div class="flex-col group_36 space-y-64">
            <div class="flex-col self-center group_8 space-y-26">
              <div class="flex-col justify-start items-center relative group_9">
                <!-- <span class="text_14 text_20">System architecture</span> -->
                <span class="text_16 pos_98">人机协同智能操作系统架构</span>
              </div>
              <span class="text_21">
                运行于传统的计算机操作系统之上，实现智能无人平台的传感器、执行器、通信系统等硬件资源的虚拟化，管理智能无人系统的感知、规划、推理等自主行为，并对人机自主协同和持续自主学习提供支撑
              </span>
            </div>
            <div class="flex-col pos_97">
              <div class="flex-col">
                <div class="flex-row group_10 items-center space-x-82">
                  <a @mouseenter="xtjg('tab1')" class="font_4" :class="{ sel: activeTab === 'tab1' }">无人平台硬件虚拟层</a>
                  <a @mouseenter="xtjg('tab2')" class="font_4" :class="{ sel: activeTab === 'tab2' }">感知层</a>
                  <a @mouseenter="xtjg('tab3')" class="font_4" :class="{ sel: activeTab === 'tab3' }">规划与推理层</a>
                  <a @mouseenter="xtjg('tab4')" class="font_4" :class="{ sel: activeTab === 'tab4' }">人机协同层</a>
                  <a @mouseenter="xtjg('tab5')" class="font_4" :class="{ sel: activeTab === 'tab5' }">虚拟空间自主学习平台</a>
                  <a @mouseenter="xtjg('tab6')" class="font_4" :class="{ sel: activeTab === 'tab6' }">作业环境模型</a>
                </div>


                <div class="flex-col justify-start items-start divider">
                  <div class="section_10"></div>
                </div>
              </div>
              <div id="tab1" class="tab visible flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-18">
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz609.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">传感器</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row self-start space-x-9">
                          <div class="flex-col justify-start text-wrapper_1">
                            <span class="text_26">激光扫描仪</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_3">
                            <span class="text_26">可见光/红外像机</span>
                          </div>
                        </div>
                        <div class="flex-row space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_1">
                            <span class="text_26">GNSS</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_3">
                            <span class="text_26">加速度计/陀螺仪</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz610.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7">实时无线通信</span>
                      <div class="flex-col">
                        <div class="flex-row group_30 space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_1">
                            <span class="text_26">自组织链路</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_3">
                            <span class="text_26">高可靠无人链路</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start self-start text-wrapper_16">
                          <span class="text_26">实时信息通信</span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz611.png" />
                    <div class="flex-col items-start group_29 space-y-12">
                      <span class="font_7">执行器</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_1">
                            <span class="text_26">姿态</span>
                          </div>
                          <div class="flex-col justify-start items-center text-wrapper_5">
                            <span class="text_26">位置</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_6">
                          <span class="text_26">速度/加速度</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_33 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图1.png" />
                    <span class="font_5 text_31">
                      无人平台硬件虚拟层对无人平台各类硬件进行虚拟化，从而屏蔽硬件的物理差异，向上提供统一的虚拟模型和访问接口，为自主协同行为提供支撑。其关键技术包括各类无人平台的传感器及执行器的虚拟模型和访问接口设计、无线通信虚拟技术及其接口等。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab1" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
              <div id="tab2" class="tab flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-18">
                  <div class="flex-col justify-start section_123">
                    <img class="imgpos_5" src="/img/ros-hmci/group668.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">感知模块</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row self-start space-x-9">
                          <div class="flex-col justify-start text-wrapper_1">
                            <span class="text_26">目标检测</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_3">
                            <span class="text_26">障碍物检测</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start self-start text-wrapper_16">
                          <span class="text_26">目标识别与跟踪</span>
                        </div>
                        <div class="flex-col justify-start self-start text-wrapper_16">
                          <span class="text_26">异构传感器数据融合框架</span>
                        </div>
                        <div class="flex-col justify-start self-start text-wrapper_16">
                          <span class="text_26">随机复杂高动态强鲁棒感知认知</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_33 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图2.png" />
                    <span class="font_5 text_31">
                      感知层不断综合各种传感器数据，基于各类感知算法，构建无人系统的作业环境模型，并确保作业环境模型总是被实时更新。其关键技术包手异构传感器融合框架以支撑多传感器数据融合技术，目标检测\识别\跟踪与障碍物检测技术、随机复杂高动态强鲁棒感知认知技术与算法等。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab2" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
              <div id="tab3" class="tab flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-18">
                  <div class="flex-col justify-start section_122">
                    <img class="imgpos_4" src="/img/ros-hmci/group7079.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">作业状态切换引擎</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row self-start space-x-9">
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">作业状态机管理</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">基于状态机的推理</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_16">
                          <span class="text_26">失效处理</span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_122">
                    <img class="imgpos_4" src="/img/ros-hmci/group7103.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7">失效处理</span>
                      <div class="flex-col">
                        <div class="flex-row group_30 space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">多目标融合规划...</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">长时多模协同混...</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_33 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图3.png" />
                    <span class="font_5 text_31">
                      规划与推理层基于感知信息和作业环境模型，按作业任务要求，开展高层次的规划和推理，将作业分解为可执行的状态机，并规划为无人平台可以执行的动作序列，由作业状态切换引擎和运动规划器组成。其关键技术包括作业状态机管理技术、基于状态机的推理技术、失效和异常处理技术等。运动规划器研究多目标融合决策与规划、长时多模协同混合控制等技术等。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab3" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
              <div id="tab4" class="tab flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-18">
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz609.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">人机互理解</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row self-start space-x-9">
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">作业指令设计</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">协同目标识别</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_16">
                          <span class="text_26">基于作业环境模型的人机交互</span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz610.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7">协同作业规划</span>
                      <div class="flex-col">
                        <div class="flex-row group_30 space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">作业预规划</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">动态作业规划</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz611.png" />
                    <div class="flex-col items-start group_29 space-y-12">
                      <span class="font_7">人机自主协同操控</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">自主协同操控...</span>
                          </div>
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">协同任务规划...</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_16">
                          <span class="text_26">协同路径规划与控制</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_33 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图4.png" />
                    <span class="font_5 text_31">
                      人机协同层支持人机自主协同，包括人机互理解模块和人机自主协同操作模块。人机互理解模块为人机自主协同提供界面、指令和基础算法，其关键技术包括作业指令设计、基于作业环境模型的人机交互以及协同目标识别等算法，协同作业规划模块实现人向无人系统的作业分配方法，作业预规划和动态作业规划等技。人机自主协同操控模块实现作业过程的人对无人系统的粗粒度控制，其关键技术包括自主协同操控通用框架、协同任务规划交互控制、协同路径规划与控制等技术。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab4" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
              <div id="tab5" class="tab flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-181">
                  <div class="flex-col justify-start section_121">
                    <img class="imgpos_3" src="/img/ros-hmci/mbz658.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">管理平台</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row self-start space-x-9">
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">资源</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">数据</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">任务</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz610.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7">虚拟学习引擎</span>
                      <div class="flex-col">
                        <div class="flex-row group_30 space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">自主强化学习</span>
                          </div>
                          <div class="flex-col justify-start text-wrapper_16">
                            <span class="text_26">人机融合学习</span>
                          </div>
                        </div>
                        <div class="flex-row group_30 space-x-9">
                          <div class="flex-col justify-start self-start text-wrapper_16">
                            <span class="text_26">大时空尺度持...</span>
                          </div>
                          <div class="flex-col justify-start self-start text-wrapper_16">
                            <span class="text_26">规模化场景学习...</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz611.png" />
                    <div class="flex-col items-start group_29 space-y-12">
                      <span class="font_7">数据生成引擎</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_6">
                            <span class="text_26">作业环境生成</span>
                          </div>
                          <div class="flex-col justify-start items-center text-wrapper_6">
                            <span class="text_26">无人系统生成</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_6">
                          <span class="text_26">虚拟目标生成</span>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_1" src="/img/ros-hmci/mbz611.png" />
                    <div class="flex-col items-start group_29 space-y-12">
                      <span class="font_7">模拟仿真引擎</span>
                      <div class="flex-col space-y-9">
                        <div class="flex-row space-x-9">
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">精准碰撞检测技术</span>
                          </div>
                          <div class="flex-col justify-start items-center text-wrapper_16">
                            <span class="text_26">全要素模拟技术</span>
                          </div>
                        </div>
                        <div class="flex-col justify-start items-center self-start text-wrapper_16">
                          <span class="text_26">并行计算仿真技术</span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_331 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图1.png" />
                    <span class="font_5 text_31">
                      虚拟空间自主学习平台运行在后台服务器集群中，用于构建虚拟空间，并实现并行可持续自主学习。其核心模块包括支持高性能、高精度、高保真、高动态的模拟仿真引擎，面向复杂环境和多型无人系统的数据生成引擎，虚拟空间自主学习引擎等。/匹配/评估技术等。针对实体空间的人机物全要素场景，研究构建全视景虚拟作业环境的智能方法和视景渲染技术。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab5" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
              <div id="tab6" class="tab flex-row items-center group_31 space-x-40">
                <div class="flex-col group_26 space-y-18">
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_2" src="/img/ros-hmci/group7310.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7 text_25">地图和障碍物</span>
                    </div>
                  </div>
                  <div class="flex-col justify-start section_12">
                    <img class="imgpos_2" src="/img/ros-hmci/group7311.png" />
                    <div class="flex-col group_29 space-y-12">
                      <span class="self-start font_7">自身位姿和速度</span>
                    </div>
                  </div>
                </div>
                <div class="flex-col group_33 space-y-56">
                  <div class="flex-col space-y-32">
                    <img class="image_17" src="/img/ros-hmci/架构图2.png" />
                    <span class="font_5 text_31">
                      作业环境模型中记录了机器和算法可理解的地图和障碍物、自身的位姿和速度等信息，以及用于识别是否关注目标的目标特征库。其关键技术包括高效的地图构建与更新技术、基于多传感器的自身位姿与速度估计、目标特征的检索/匹配/评估技术等。
                    </span>
                  </div>
                  <!-- <div class="flex-row self-start section_8 view_4 space-x-4">
                    <a :href="links.os_tab6" class="font_7 text_33">查看详情</a>
                    <img class="shrink-0 image_18" src="/img/ros-hmci/样式.png" />
                  </div> -->
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start items-center section_13">
          <div id="osDownload" class="flex-col space-y-48">
            <div class="flex-col justify-start items-center self-center relative group_12">
              <!-- <img class="image_19" src="/img/ros-hmci/a5163f60f9fa71fbeec94b26f590f7a1.png" /> -->
              <span class="font_3 text_14 text_34 pos_17">系统下载</span>
            </div>

            <div class="flex-row space-x-14">

              <div class="flex-row relative">
                <div class="flex-col section_14">
                  <div v-for="(repo, index) in repos" :id="`sec${index + 1}`" :ref="`sec${index + 1}`" class="sec repo"
                    @mouseenter="xtxz(`sec${index + 1}`); selectedRepoName = repo.Name; selectedRepoAlias = repo.Alias; selectedRepoOwner = repo.OwnerName">
                    <div class="flex-col items-start section_16 xz_active"
                      :style="{ display: index === 0 ? '' : 'none' }">
                      <span class="font_10">人机协同智能操作系统</span>
                      <div class="flex-col justify-start text-wrapper_7">
                        <span class="font_2 text_36">{{ repo.Alias }}</span>
                      </div>
                    </div>
                    <div class="flex-col justify-start items-center section_31 xz_normal"
                      :style="{ display: index === 0 ? 'none' : '' }">
                      <span class="text_37">人机协同智能操作系统</span>
                      <span class="text_39">-{{ repo.Alias }}</span>
                    </div>
                  </div>
                </div>
              </div>

              <div class="flex-col group_34">
                <div ref="dataset_info" class="flex-col section_17 visible">
                  <span class="self-start font_10 text_35">人机协同智能操作系统-{{ selectedRepoAlias }}简介</span>
                  <div class="flex-col items-start group_13 space-y-2">
                    <span class="font_11"> {{ description }} </span>
                  </div>
                  <div class="flex-row self-start section_18 space-x-2">
                    <a :href="`/${selectedRepoOwner}/${selectedRepoName}/datasets`" class="font_8 text_38">更多详情</a>
                    <img class="shrink-0 image_21" src="/img/ros-hmci/8aacfa799145b95562e808dcf6004f2d.png" />
                  </div>

                  <div class="flex-col relative section_19">

                    <div class="flex-col justify-start items-center text-wrapper_8 pos_18">
                      <span class="font_2">版本列表</span>
                    </div>

                    <div v-for="attachment in attachments">
                      <div class="flex-col">
                        <div class="group_14 view_5">
                          <span class="font_14">{{ attachment.Name }}</span>
                          <a :href="`/attachments/${attachment.UUID}?type=0`"><img class="shrink-0 image_22"
                              src="/img/ros-hmci/mbz614.png" /></a>
                        </div>
                        <div class="flex-row items-baseline group_14">
                          <a :href="`/${selectedRepoOwner}/${selectedRepoName}/datasets`"
                            class="font_13 text_42">本次改动内容...</a>
                          <span class="font_15">{{ attachment.CreatedUnix | formatDate }}</span>
                        </div>
                      </div>

                      <div class="divider_2"></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppBanner from "../components/AppBanner.vue";
import linksConfig from './links.js'
import axios from 'axios';

export default {
  name: "OSSystem",
  data() {
    return {
      activeTab: 'tab1',
      links: linksConfig,
      repos: [],
      description: '',
      attachments: [],
      selectedRepoAlias: '',
      selectedRepoName: '',
      selectedRepoOwner: '',
      showLeft: true,
      showRight: true,
    };
  },
  filters: {
    formatDate(timestamp) {
      const date = new Date(timestamp * 1000);
      const year = date.getFullYear();
      const month = ('0' + (date.getMonth() + 1)).slice(-2);
      const day = ('0' + date.getDate()).slice(-2);
      return `${year}-${month}-${day}`;
    }
  },
  components: {
    AppBanner,
  },
  created() {
    this.fetchRepos();
  },
  mounted() {
    // window.addEventListener('scroll', this.handleScroll);
    // this.xtxz("sec1");
    this.toLocal();
  },
  beforeDestroy() {
    // window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll() {
      const leftElem = this.$refs.leftElem;
      const rightElem = this.$refs.rightElem;

      const windowHeight = window.innerHeight;
      const scrollY = window.scrollY;

      const leftElemTop = leftElem.getBoundingClientRect().top + scrollY;
      const rightElemTop = rightElem.getBoundingClientRect().top + scrollY;

      const threshold = windowHeight * 0.5;

      if (scrollY + windowHeight > leftElemTop + threshold) {
        this.showLeft = true;
      } else {
        this.showLeft = false;
      }

      if (scrollY + windowHeight > rightElemTop + threshold) {
        this.showRight = true;
      } else {
        this.showRight = false;
      }
    },

    async fetchRepos() {
      try {
        const apiUrl = '/explore/repos/search?q=&topic=ros-hmci-os-xjykd&sort=mostpopular&pageSize=15&page=1';
        const response = await axios.get(apiUrl);

        // 如果返回值的Code为0，表示请求成功，将Repos数据赋值给data属性repos
        if (response.data.Code === 0) {
          this.repos = response.data.Data.Repos;

          // 遍历仓库列表并调用fetchCurrentRepoData方法
          for (const repo of this.repos) {
            await this.fetchCurrentRepoData(repo.OwnerName, repo.Name);
          }

        } else {
          // 处理请求失败的情况
          console.error('请求失败：', response.data.Msg);
        }
      } catch (error) {
        // 处理请求异常的情况
        console.error('请求异常：', error);
      }
    },
    async fetchCurrentRepoData(repoOwnerName, repoName) {
      try {
        const apiUrl = `/${repoOwnerName}/${repoName}/datasets/current_repo_m?q=&page=1&q=&type=0`;
        const response = await axios.get(apiUrl);

        // 处理请求成功的情况
        if (response.data.result_code === "0") {
          // const test1 = response.data.data;
          // console.log(test1);

          const responseData = JSON.parse(response.data.data); // 解析 "data" 字段的 JSON 字符串
          // console.log(responseData);

          const data = responseData[0]; // 假设 data 数组中的第一个对象包含所需数据
          this.description = data.Description;
          this.attachments = data.Attachments;
        } else {
          // 处理请求失败的情况
          console.error("请求失败：", response.data.Msg);
        }
      } catch (error) {
        // 处理请求异常的情况
        console.error("请求异常：", error);
      }
    },

    toLocal() {
      // console.log(this.$route);
      // 查找存储的锚点id
      let Id = this.$route.query.id;

      let toElement = document.getElementById(Id);
      // console.log(toElement, "toElement");
      //锚点存在跳转
      if (Id) {
        toElement.scrollIntoView();
      }
    },
    xtjg(tabID) {
      const tabs = document.querySelectorAll(".tab");
      const tabbar = document.querySelector(".section_10");

      tabs.forEach((tab) => {
        if (tab.id == tabID) {
          tab.classList.add("visible");
        } else {
          tab.classList.remove("visible");
        }
      });

      const tabLefts = {
        "tab1": "0px",
        "tab2": "205px",
        "tab3": "385px",
        "tab4": "575px",
        "tab5": "815px",
        "tab6": "1029px",
      };

      tabbar.style.left = tabLefts[tabID] || "0px";

      // 更新激活的标签
      this.activeTab = tabID;
    },

    xtxz(secID) {
      const secs = document.querySelectorAll(".sec");

      secs.forEach((sec) => {
        const active = sec.querySelector(".xz_active");
        const normal = sec.querySelector(".xz_normal");

        if (sec.id == secID) {
          active.style.display = "";
          normal.style.display = "none";
        } else {
          active.style.display = "none";
          normal.style.display = "";
        }
      });


      const index = parseInt(secID.replace("sec", "")) - 1;
      this.currentRepo = this.repos[index];
    },
    destroyed() {
      localStorage.setItem("toId", "");
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.slide-in-left {
  transform: translateX(-100%);
  opacity: 0;
  transition: transform 0.5s ease-out, opacity 0.5s ease-out;
}

.slide-in-right {
  transform: translateX(100%);
  opacity: 0;
  transition: transform 0.5s ease-out, opacity 0.5s ease-out;
}

.slide-in-left.show,
.slide-in-right.show {
  transform: translateX(0);
  opacity: 1;
}

.page {
  /* background-color: #0b0e1d; */
  width: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  height: 100%;
  margin-bottom: -80px;
}

.group_25 {
  position: relative;
  top: 0px;
  background: url(/img/ros-hmci/mbz605.png);
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  height: 600px;
  width: 100%;
}

.section_5 {
  padding: 10px 360px;
  background-color: #757c8d2e;
}

.space-x-14>*:not(:first-child) {
  margin-left: 14px;
}

.image_10 {
  width: 33px;
  height: 21px;
}

.font_2 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  line-height: 13px;
  color: #ffffff;
}

.section_6 {
  padding: 80px 0 149px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  box-shadow: 0px 3px 12px #0000000a;
}

.space-y-72>*:not(:first-child) {
  margin-top: 72px;
}

.group_3 {
  padding-top: 0px;
}

.text_16 {
  color: #181818;
  font-size: 36px;
  font-family: Alibaba PuHuiTi;
  line-height: 33.5px;
}

.font_3 {
  font-size: 50px;
  font-family: PingFang SC;
  line-height: 37px;
}

.text_14 {
  /* opacity: 0.11; */
}

.text_15 {
  color: #24a19b;
}

.pos_4 {
  position: relative;
  top: -48px;
}

.space-x-92>*:not(:first-child) {
  margin-left: 92px;
}

.group_5 {
  margin-top: 23px;
}

.space-y-54>*:not(:first-child) {
  margin-top: 56px;
}

.space-x-16>*:not(:first-child) {
  margin-left: 16px;
}

.image_12 {
  width: 21px;
  height: 19px;
}

.font_41 {
  font-size: 18px;
  font-family: Alibaba PuHuiTi;
  line-height: 17px;
  color: #181818;
}

.text_17 {
  color: #323232;
  line-height: 16.5px;
}

.group_6 {
  width: 562px;
  margin-top: 22px;
}

.font_5 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 30px;
  color: #4d4d4d;
}

.font_51 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #757c8d;
}

.group_7 {
  margin-top: 30px;
}

.space-x-8>*:not(:first-child) {
  margin-left: 8px;
}

.text-wrapper_2 {
  padding: 8px 6px;
  height: 28px;
  background-color: rgba(229, 179, 114, 0.16);
  border-radius: 4px;
}

.info {
  position: absolute;
  top: 104px;
  left: 18%;
}

.info>*:not(:first-child) {
  margin-top: 36px;
}

.info-div {
  width: 486.09px;
  height: 53px;
  background-image: linear-gradient(150.44deg, #00e850 0%, #0100be 100%);
  border-radius: 28px;
}

.info1 {
  font-family: sucaijishikufangti;
  color: #ffffff;
  font-size: 45px;
  letter-spacing: 7px;
}

.info2 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 21px;
  letter-spacing: 1px;
}

.info3 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 40px;
  letter-spacing: 8px;
}

.font_6 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  line-height: 13px;
  color: #bb7f41;
}

.section_8 {
  padding: 12px 18px;
  background-image: linear-gradient(152deg, #35cf67 18.9%, #179d8c 69.8%);
  border-radius: 4px;
}

.text_19 {
  color: #ffffff;
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 15px;
}

.image_14 {
  margin-left: 4px;
  width: 17px;
  height: 17px;
}

.image_27 {
  width: 559px;
  height: 344px;
}

.section_9 {
  background-color: #ffffff00;
}

.pos_9 {
  position: absolute;
}

.pos_98 {
  position: relative;
}

.pos_99 {
  position: relative;
  top: 6px;
}

.section_30 {
  background-color: #f6f9fb;
  box-shadow: 0px 3px 12px #0000000a;
}

.image_1 {
  width: 100%;
  height: 965px;
}

.image_13 {
  width: 100%;
  height: 1323px;
}

.pos_1 {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
}

.group_1 {
  padding-top: 30px;
  padding-bottom: 60px;
}

.group_36 {
  position: relative;
  top: 48px;
}

.space-y-64>*:not(:first-child) {
  margin-top: 64px;
}

.group_8 {
  width: 880px;
}

.space-y-26>*:not(:first-child) {
  margin-top: 40px;
}

.text_20 {
  color: #24a19b;
  font-size: 50px;
  font-family: PingFang SC;
  line-height: 46.5px;
}

.image_15 {
  width: 420px;
  height: 34px;
}

.pos_6 {
  position: absolute;
  left: 50%;
  bottom: 0;
  transform: translateX(-50%);
}

.text_21 {
  width: 885px;
  text-align: center;
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #757c8d;
}

.group_10 {
  padding: 0 6px 16px;
}

.space-x-82>*:not(:first-child) {
  margin-left: 100px;
}

.text_24 {
  line-height: 16.5px;
}

.divider {
  width: 1200px;
  height: 0px;
  border: 1px dashed #000;
  border-color: #caced8;
}

.section_10 {
  position: relative;
  left: 0px;
  bottom: 2px;
  background-image: linear-gradient(90.2deg, #1ad1c3 8.9%, #02e588 81.9%);
  border-radius: 15px;
  width: 170px;
  height: 3px;
  transition: all 0.3s;
}

.group_31 {
  padding-top: 0px;
  padding-bottom: -5px;
}

.space-x-40>*:not(:first-child) {
  margin-left: 40px;
}

.group_26 {
  width: 310px;
}

.space-y-18>*:not(:first-child) {
  margin-top: 18px;
}

.space-y-181>*:not(:first-child) {
  margin-top: 9px;
}

.section_12 {
  position: relative;
  width: 310px;
  height: 130px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 3px 10px rgba(164, 169, 181, 0.35);
}

.section_121 {
  position: relative;
  width: 310px;
  height: 91px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 3px 10px rgba(164, 169, 181, 0.35);
}

.section_122 {
  position: relative;
  width: 310px;
  height: 197px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 3px 10px rgba(164, 169, 181, 0.35);
}

.section_123 {
  position: relative;
  width: 310px;
  height: 411px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 3px 10px rgba(164, 169, 181, 0.35);
}

.group_29 {
  padding: 16px 16px;
  width: 310px;
  z-index: 1;
}

.imgpos_1 {
  position: absolute;
  right: 0px;
  top: 14px;
  height: 123px;
  width: 123px;
  z-index: 0;
}

.imgpos_2 {
  position: absolute;
  right: 0px;
  height: 100%;
  width: 100%;
  z-index: 0;
}

.imgpos_3 {
  position: absolute;
  right: 16px;
  top: 10px;
  width: 33%;
  z-index: 0;
}

.imgpos_4 {
  position: absolute;
  right: 20px;
  bottom: 0;
  width: 42%;
  z-index: 0;
}

.imgpos_5 {
  position: absolute;
  right: 0px;
  bottom: 0;
  z-index: 0;
  height: 286px;
}

.font_7 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  color: #323232;
}

.text_25 {
  line-height: 15px;
}

.space-y-9>*:not(:first-child) {
  margin-top: 9px;
}

.text-wrapper_3 {
  padding: 8px 10px;
  background-color: #a4a9b529;
  border-radius: 4px;
  height: 28px;
  /* width           : 128px; */
}

.font_8 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  line-height: 13px;
  color: #4d4d4d;
}

.text_26 {
  margin: 0 auto;
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 14px;
  text-align: center;
  line-height: 13px;
}

.text_27 {
  margin-left: 12px;
  margin-right: 8px;
}

.text-wrapper_1 {
  padding: 8px 0px;
  background-color: #a4a9b529;
  border-radius: 4px;
  width: 88px;
  height: 28px;
}

.font_9 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  line-height: 10px;
  color: #4d4d4d;
}

.text_28 {
  margin: 0 10px;
}

.group_28 {
  padding: 20px 16px 17px;
  width: 310px;
  z-index: 1;
}

.group_30 {
  padding-bottom: 9px;
}

.space-x-9>*:not(:first-child) {
  margin-left: 9px;
}

.text_29 {
  margin-left: 14px;
  margin-right: 10px;
}

.text-wrapper_16 {
  padding: 8px 10px;
  background-color: #a4a9b529;
  border-radius: 4px;
}

.text_30 {
  margin-left: 10px;
  margin-right: 6px;
}

.section_1 {
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 3px 10px #a4a9b559;
  height: 130px;
}

.group_27 {
  padding: 20px 16px;
  width: 310px;
}

.text_4 {
  margin-left: 48px;
  margin-top: 4px;
  font-size: 15px;
  line-height: 14px;
}

.text-wrapper_5 {
  padding: 8px 0;
  background-color: #a4a9b529;
  border-radius: 4px;
  width: 68px;
  height: 28px;
}

.tab {
  position: absolute;
  width: 100%;
  top: 288px;
  opacity: 0;
  transition: opacity 0.3s ease-in-out;
  pointer-events: none;
}

.tab.visible {
  opacity: 1;
  pointer-events: all;
}

.font_4 {
  font-size: 18px;
  font-family: Alibaba PuHuiTi;
  line-height: 17px;
  color: #181818;
}

.font_4.sel {
  color: #24a19b;
}

.pos_97 {
  height: 583px;
}

.text_32 {
  line-height: 12.5px;
}

.text-wrapper_6 {
  padding: 8px 0;
  background-color: #a4a9b529;
  border-radius: 4px;
  width: 102px;
}

.group_33 {
  width: 852.5px;
}

.group_331 {
  width: 852.5px;
  position: relative;
  top: -41px;
}

.space-y-56>*:not(:first-child) {
  margin-top: 48px;
}

.space-y-32>*:not(:first-child) {
  margin-top: 30px;
}

.image_17 {
  margin-right: 0px;
  width: 850px;
  height: 239px;
}

.text_31 {
  margin-left: 3px;
}

.view_4 {
  margin-left: 3px;
}

.text_33 {
  color: #ffffff;
}

.image_18 {
  margin: auto;
  width: 11px;
}

.section_13 {
  margin-top: -13px;
  padding: 130px 0 154px;
  background-image: url("/img/ros-hmci/mbz612.png");
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
}

.space-y-48>*:not(:first-child) {
  margin-top: 48px;
}

.group_12 {
  padding-top: 18px;
  width: 228px;
}

.image_19 {
  width: 138px;
  height: 33px;
}

.text_34 {
  font-size: 36px;
  color: #181818;
}

.pos_17 {
  position: absolute;
  left: 0;
  right: 0;
  top: -45px;
}

.section_14 {
  padding-bottom: 156px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow:0px 0px 12px rgba(164, 169, 181, 0.24);
  width: 280px;
  height: 450px;
}

.section_15 {
  background-image: linear-gradient(168.36deg, #24a19b 0%, #19bf77 64.7%);
  border-radius: 4px;
}

.section_15:hover {
  transition: opacity 0.7s ease-in-out;
}

.section_16 {
  padding: 14px 20px;
  background-image: url("/img/ros-hmci/group3576.png");
  border-radius: 4px;
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 293px;
}

.section_16>*:not(:first-child) {
  margin-top: 8px;
}

.font_10 {
  font-size: 16px;
  font-family: PingFang SC;
  /* line-height: 15px; */
  color: #ffffff;
}

.text-wrapper_7 {
  padding: 7px 12px;
  border-radius: 4px;
  border: solid 1px #ffffff;
}

.section_31 {
  height: 87px;
}

.font_12 {
  font-size: 14px;
  font-family: PingFang SC;

  color: #323232;
}

.text_37 {
  font-size: 14px;
  font-family: PingFang SC;
  color: #323232;
  align-self: flex-start;
  margin-left: 20px;
  margin-top: 26px;
}

.font_13 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  color: #757c8d;
}

.text_39 {
  align-self: flex-start;
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  color: #757c8d;
  margin-left: 20px;
  margin-top: 6px;
}

.text_40 {
  margin-left: 20px;
  margin-top: 34px;
}

.text_41 {
  margin-left: 20px;
  margin-top: 16px;
  line-height: 10px;
}

.text_44 {
  margin-left: 20px;
  margin-top: 36px;
}

.text_45 {
  margin-left: 22px;
  margin-top: 16px;
}

.text_48 {
  margin-left: 20px;
  margin-top: 34px;
}

.text_49 {
  margin-left: 22px;
  margin-top: 16px;
  line-height: 13.5px;
}

.image_20 {
  position: relative;
  right: 1px;
  top: 30px;
  border-radius: 1px;
  width: 17px;
  height: 28px;
}

.group_34 {
  width: 893px;
}

.space-y-12>*:not(:first-child) {
  margin-top: 14px;
}

.section_17 {
  position: absolute;
  width: 884px;
  padding: 22px 22px 72px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow:0px 0px 12px rgba(164, 169, 181, 0.24);
  margin-left: 13px;
  opacity: 0;
  transition: opacity 0.7s ease-in-out;
  pointer-events: none;
}

.section_17.visible {
  opacity: 1;
  pointer-events: all;
}

.text_35 {
  color: #181818;
}

.group_13 {
  margin-top: 22px;
}

.space-y-2>*:not(:first-child) {
  margin-top: 2px;
}

.font_11 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  line-height: 17px;
  color: #4d4d4d;
}

.section_18 {
  margin-top: 16px;
  padding: 6px;
  border-radius: 2px;
  border: solid 1px #24a19b;
}

.space-x-2>*:not(:first-child) {
  margin-left: 2px;
}

.text_38 {
  color: #24a19b;
}

.image_21 {
  width: 8px;
  height: 12px;
}

.section_19 {
  margin-top: 19px;
  padding: 35px 8px 18px;
  background-image:linear-gradient(180deg,#eff6f3 0%,#ffffff 100%);
  border-radius: 4px;
  box-shadow:0px 3px 12px rgba(0, 0, 0, 0.04);
  border: solid 1px #ffffff;
  min-height: 252px;
}

.group_14 {
  display: flex;
  justify-content: space-between;
  margin-top: 12px;
  padding: 0 4px;
}

.view_5 {
  margin-top: 0px;
}

.space-x-504>*:not(:first-child) {
  margin-right: 0px;
}

.font_14 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
  /* line-height: 13px; */
  color: #323232;
}

.image_22 {
  width: 17px;
  height: 17px;
}

.space-x-660>*:not(:first-child) {
  margin-left: 660px;
}

.text_42 {
  font-size: 13px;
  /* line-height: 12px; */
}

.font_15 {
  font-size: 12px;
  font-family: Alibaba PuHuiTi;
  line-height: 8.5px;
  color: #757c8d;
}

.divider_2 {
  margin: 14px 0px;
  border: 1px dashed #fff;
  border-color: #caced8;
}

.group_15 {
  margin-top: 14px;
}

.space-x-512>*:not(:first-child) {
  margin-left: 512px;
}

.text_46 {
  font-size: 13px;
  line-height: 12px;
}

.group_16 {
  margin-top: 14px;
  padding: 0 6px;
}

.space-x-550>*:not(:first-child) {
  margin-left: 550px;
}

.text_50 {
  font-size: 13px;
  line-height: 12px;
}

.text-wrapper_8 {
  padding: 6px 0;
  background-color: #24a19b;
  border-radius: 4px 0px 0px 0px;
  width: 96px;
}

.pos_18 {
  position: absolute;
  left: 0;
  top: 0;
}

.section_20 {
  padding: 22px 22px 30px;
  background-image: linear-gradient(180deg, #f6f9fb 0%, #ffffff 100%);
  border-radius: 4px;
  box-shadow: 0px 0px 12px #0716156b;
}

.space-y-20>*:not(:first-child) {
  margin-top: 20px;
}

.text_52 {
  color: #323232;
}

.space-x-12>*:not(:first-child) {
  margin-left: 12px;
}

.text-wrapper_9 {
  padding: 6px 0;
  background-color: #caced836;
  border-radius: 2px;
  height: 26px;
}

.text_53 {
  margin-left: 12px;
  margin-right: 4px;
}

.text-wrapper_10 {
  padding: 10px 0 6px;
  background-color: #caced836;
  border-radius: 2px;
  width: 52px;
  height: 26px;
}

.text_56 {
  line-height: 10.5px;
}

.text-wrapper_11 {
  padding: 6px 0;
  background-color: #caced836;
  border-radius: 2px;
  width: 72px;
  height: 26px;
}

.text_54 {
  line-height: 13.5px;
}

.text_55 {
  line-height: 11px;
}

.text-wrapper_12 {
  padding: 8px 0 6px;
  background-color: #caced836;
  border-radius: 2px;
  width: 72px;
  height: 26px;
}

.text_57 {
  line-height: 12px;
}

.text-wrapper_13 {
  padding: 6px 0;
  background-color: #caced836;
  border-radius: 2px;
  width: 84px;
  height: 26px;
}

.text_58 {
  line-height: 13.5px;
}

.section_21 {
  padding: 4px 8px;
  border-radius: 4px;
  height: 26px;
  border: solid 1px #caced8;
}

.space-x-4>*:not(:first-child) {
  margin-left: 6px;
}

.image_24 {
  width: 9px;
  height: 6px;
}

.section_22 {
  background-color: #1b2440;
}

.group_23 {
  padding: 65px 0 62px;
}

.image-wrapper_3 {
  padding: 22px 0 14px;
  background-color: #ffffff;
  border-radius: 13px;
  width: 100px;
}

.image_25 {
  width: 91px;
  height: 64px;
}

.text_1 {
  margin-top: 32px;
}

.text_63 {
  margin-top: 26px;
  color: #bdc2d1;
  line-height: 14.5px;
}

.group_38 {
  margin-left: 195px;
  margin-top: 4px;
}

.space-x-148>*:not(:first-child) {
  margin-left: 148px;
}

.space-y-23>*:not(:first-child) {
  margin-top: 23px;
}

.font_16 {
  font-size: 18px;
  font-family: 苹方;
  line-height: 17px;
  color: #ffffff;
}

.text_59 {
  line-height: 16.5px;
}

.font_17 {
  font-size: 14px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #bdc2d1;
}

.text_60 {
  line-height: 16.5px;
}

.group_37 {
  margin-left: 218px;
}

.space-y-24>*:not(:first-child) {
  margin-top: 24px;
}

.text_61 {
  line-height: 16.5px;
}

.space-x-76>*:not(:first-child) {
  margin-left: 76px;
}

.image_26 {
  border-radius: 4px;
  width: 99px;
  height: 100px;
}

.text_62 {
  line-height: 12.5px;
}

.section_23 {
  padding: 29px 0;
  background-color: #222a42;
}

.space-y-6>*:not(:first-child) {
  margin-top: 6px;
}

.group_21 {
  width: 395px;
  height: 12.5px;
}

.font_18 {
  font-size: 12px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #bdc2d1;
}

.text_64 {
  line-height: 12.5px;
}

.text_65 {
  margin-top: -12px;
  line-height: 12.5px;
}

.text_66 {
  line-height: 14px;
  text-align: center;
  width: 381px;
}
</style>
