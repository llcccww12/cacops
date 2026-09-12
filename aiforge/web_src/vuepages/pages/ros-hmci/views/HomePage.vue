<template>
  <div class="flex-col relative page" style="padding-bottom: 0px">
    <div class="flex-col section_6">
      <div class="section_6_cont">
        <div class="homeHeadCont">
          <div class="flexCont">
            <span>欢迎来到</span>
            <span class="homeBox1">智能学习</span>
          </div>
          <div>
            <span class="homeBox2">人机协同</span>
            <span>智能操作系统</span>
          </div>
          <div class="homeBox3">AI学习者的一体化智能学习交流平台</div>
          <a class="font_5 gotoDownload" @click="getdownload('/os-system', 'osDownload')">立即下载</a>
        </div>

        <!-- 右边动效区域 -->
        <div class="banner_right">
          <!-- <img class="sec_1" src="/img/ros-hmci/group7392.png">
          <img class="sec_2" src="/img/ros-hmci/group7391.png">
          <img class="sec_3" src="/img/ros-hmci/group7390.png">
          <img class="sec_4" src="/img/ros-hmci/group7389.png">
          <img class="banner_ros bounce" src="/img/ros-hmci/banner-ros.png">
          <img class="banner_bg" src="/img/ros-hmci/首页路径.png"> -->
          <img class="homepage1" src="/img/ros-hmci/homepage1.png">
          <img class="homepage2" src="/img/ros-hmci/homepage2.png">
          <img class="homepage3" src="/img/ros-hmci/homepage3.png">
          <img class="homepage4" src="/img/ros-hmci/homepage4.png">
        </div>

        <!-- <div class="flex-col self-start group_5">
          <span class="self-start font_4 text_15">欢迎来到</span>
          <span class="homeBox1">智能学习</span>
          <div class="flex-row group_6 space-x-8">
            <img class="shrink-0 image_11" src="/img/ros-hmci/c1b8b11fc5718ac77fbc93056561e611.png" />
            <span>人机协同</span>
            <span class="font_4">智能操作系统</span>
          </div>
          <div class="info-div">
            <span class="text_16">AI学习者的一体化智能学习交流平台</span>
          </div>
        </div> -->
        <NavigationBar></NavigationBar>
      </div>
    </div>

    <div class="flex-col justify-start items-center section_10 pos_3">
      <div>
        <span class="font_8 pos_80">人机协同智能操作系统简介</span>
      </div>
      <div class="flex-row section_11 space-x-40 pos_5">
        <div class="flex-auto self-start group_13 ">
          <span class="font_9 text_28">”</span>
          <span class="text_27">
            瞄准人机协同自主作业需求，针对“可持续自主学习型智能操作系统”和“人机物融合、多主体协同计算平台”两大科学问题，着眼构建新一代人工智能基础软件，创新和突破相关基础理论和关键技术，研制人机协同智能操作系统原型版本，构建支持操作系统开源生态构建和智能资源共享的开源开放基础平台，面向典型场景和任务开展应用验证。
          </span>
        </div>
        <img class="shrink-0 self-center image_17" src="/img/ros-hmci/mbz549.png" />
      </div>
    </div>

    <div>
      <div class="flex-row items-center group_14 pos_7">
        <div class="flex-col group_14 space-y-10" @mouseenter="() => mouseEnter(0)">
          <img class="image_18" :class="{ 'bubble-animation': hoveredImage === 0 }" src="/img/ros-hmci/作业环境.png" />
          <a class="text_31" :style="{ color: activeIndex === 0 ? '#24a19b' : '' }">架构设计与系统集成</a>
        </div>
        <div class="flex-col group_15 space-y-10" @mouseenter="() => mouseEnter(1)">
          <img class="self-center image_18" :class="{ 'bubble-animation': hoveredImage === 1 }"
            src="/img/ros-hmci/人机协同体系架构.png" />
          <a class="text_31" :style="{ color: activeIndex === 1 ? '#24a19b' : '' }">作业规划器与示范应用</a>
        </div>
        <div class="flex-col group_15 space-y-10" @mouseenter="() => mouseEnter(2)">
          <img class="image_18" :class="{ 'bubble-animation': hoveredImage === 2 }" src="/img/ros-hmci/持续自主学习.png" />
          <a class="text_31" :style="{ color: activeIndex === 2 ? '#24a19b' : '' }">作业规划器与示范应用</a>
        </div>
        <div class="flex-col group_17 space-y-10" @mouseenter="() => mouseEnter(3)">
          <img class="image_18" :class="{ 'bubble-animation': hoveredImage === 3 }" src="/img/ros-hmci/复杂环境建模.png" />
          <a class="text_31" :style="{ color: activeIndex === 3 ? '#24a19b' : '' }">作业环境仿真器与数字孪生平台</a>
        </div>
        <div class="tabBar"></div>
      </div>

      <div class="flex-col justify-start items-end group_20 pos_9" id="tab1" style="display: ''">
        <div class="flex-col relative section_17">
          <div class="flex-row group_21">
            <span class="self-center font_11 text_36">
              针对复杂应用环境，开展人机自主协同架构设计，突破无人平台硬件虚拟层设计、适应环境的感知/规划/推理、人机自主协同操控、面向作业的实时无线通信等关键技术，构建人机自主协同操控子系统，研制支持多种异构智能硬件/软件/数据资源和人机互理解的智能操作系统原型。
            </span>
            <span class="self-start font_9 text_35">”</span>
          </div>
          <div class="flex-col items-start relative group_22">
            <div class="flex-col justify-start items-center text-wrapper_3">
              <span class="font_5 text_37">面向开源</span>
            </div>
            <!-- <div class="flex-col justify-start items-center text-wrapper_5">
              <a :href="links.home_tab1" class="font_5 text_39">了解更多</a>
            </div> -->
            <div class="flex-col justify-start items-center group_23 pos_11">
              <img class="image_26" src="/img/ros-hmci/白色背景.png" />
              <div class="flex-col justify-start items-center text-wrapper_4 pos_12">
                <span class="text_38">质量保证机制</span>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start section_15 pos_10">
          <div class="flex-col tab_sec1 space-y-134">
          </div>
        </div>
      </div>
      <div class="flex-col justify-start items-end group_20 pos_9" id="tab2" style="display: none">
        <div class="flex-col relative section_17">
          <div class="flex-row group_21">
            <span class="self-center font_11 text_36">
              针对复杂环境下无人系统作业的机器可理解可执行需求，聚焦无人系统作业规划与作业调度问题，突破作业环境模型构建、协同作业规划、作业状态切换引擎、面向作业的人机互理解等关键技术，设计实现适应复杂环境和支持复杂作业的作业规划器子系统，开展面向反恐侦察和农田作业的人机协同智能操作系统典型应用示范。
            </span>
            <span class="self-start font_9 text_35">”</span>
          </div>
          <div class="flex-col items-start relative group_22">
            <div class="flex-col justify-start items-center text-wrapper_3">
              <span class="font_5 text_37">面向开源</span>
            </div>
            <!-- <div class="flex-col justify-start items-center text-wrapper_5">
              <a :href="links.home_tab2" class="font_5 text_39">了解更多</a>
            </div> -->
            <div class="flex-col justify-start items-center group_23 pos_11">
              <img class="image_26" src="/img/ros-hmci/白色背景.png" />
              <div class="flex-col justify-start items-center text-wrapper_4 pos_12">
                <span class="text_38">质量保证机制</span>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start section_15 pos_10">
          <div class="flex-col tab_sec2 space-y-134">
          </div>
        </div>
      </div>
      <div class="flex-col justify-start items-end group_20 pos_9" id="tab3" style="display: none">
        <div class="flex-col relative section_17">
          <div class="flex-row group_21">
            <span class="self-center font_11 text_36">
              针对人机协同智能系统快速赋能和持续增能的需求，突破智能体依托虚拟环境自主学习行为策略的核心机理，构造支撑虚拟空间内自主持续学习和利用人类智能数据进行学习的虚拟学习引擎；突破操作系统开源生态构建和演化的核心使能机理，以及运行时知识、模型、算法等智能资源的规模化共享机制，搭建面向人机协同智能操作系统软件生态链构造和运行时后端支撑的开源开放基础平台。
            </span>
            <span class="self-start font_9 text_35">”</span>
          </div>
          <div class="flex-col items-start relative group_22">
            <div class="flex-col justify-start items-center text-wrapper_3">
              <span class="font_5 text_37">面向开源</span>
            </div>
            <!-- <div class="flex-col justify-start items-center text-wrapper_5">
              <a :href="links.home_tab3" class="font_5 text_39">了解更多</a>
            </div> -->
            <div class="flex-col justify-start items-center group_23 pos_11">
              <img class="image_26" src="/img/ros-hmci/白色背景.png" />
              <div class="flex-col justify-start items-center text-wrapper_4 pos_12">
                <span class="text_38">质量保证机制</span>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start section_15 pos_10">
          <div class="flex-col tab_sec3 space-y-134">
          </div>
        </div>
      </div>
      <div class="flex-col justify-start items-end group_20 pos_9" id="tab4" style="display: none">
        <div class="flex-col relative section_17">
          <div class="flex-row group_21">
            <span class="self-center font_11 text_36">
              针对虚拟空间自主学习对作业环境仿真的需求与挑战，开展跨传统计算机操作系统的并行计算架构设计，研究人机协同仿真技术，攻克复杂作业环境智能建模、视景渲染、物理过程仿真技术，构建具有“高可扩展、高逼真、高保真、高实时”的作业环境仿真器与数字孪生平台，为可持续自主学习提供大数据支撑，以课题1和课题2的应用需求为牵引开展数字孪生平台的验证和优化。
            </span>
            <span class="self-start font_9 text_35">”</span>
          </div>
          <div class="flex-col items-start relative group_22">
            <div class="flex-col justify-start items-center text-wrapper_3">
              <span class="font_5 text_37">面向开源</span>
            </div>
            <!-- <div class="flex-col justify-start items-center text-wrapper_5">
              <a :href="links.home_tab4" class="font_5 text_39">了解更多</a>
            </div> -->
            <div class="flex-col justify-start items-center group_23 pos_11">
              <img class="image_26" src="/img/ros-hmci/白色背景.png" />
              <div class="flex-col justify-start items-center text-wrapper_4 pos_12">
                <span class="text_38">质量保证机制</span>
              </div>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start section_15 pos_10">
          <div class="flex-col tab_sec4 space-y-134">
          </div>
        </div>
      </div>
    </div>


    <div class="flex-col justify-start items-center pos_89 ">
      <span class="font_8 pos_80">项目协同</span>
      <div class="xmxt justify-center items-center flex-row">
        <div class="justify-center items-center flex-col sjx">
          <img class="image_42" src="/img/ros-hmci/mbz545.png" />
          <a class="text_73">{{ systemComponentTotal }}</a>
          <a class="text_74">系统组件</a>
        </div>
        <div class="justify-center items-center flex-col sjx">
          <img class="image_42" src="/img/ros-hmci/mbz546.png" />
          <a class="text_73">{{ applicationSoftwareTotal }}</a>
          <a class="text_74">应用软件</a>
        </div>
        <div class="justify-center items-center flex-col sjx">
          <img class="image_42" src="/img/ros-hmci/mbz547.png" />
          <a class="text_73">{{ statisticsData.developer }}</a>
          <a class="text_74">开发者</a>
        </div>
        <div class="justify-center items-center flex-col sjx">
          <img class="image_42" src="/img/ros-hmci/mbz548.png" />
          <a class="text_73">{{ statisticsData.organization }}</a>
          <a class="text_74">组织</a>
        </div>
      </div>
    </div>

    <div class="flex-col justify-start items-center section_19 pos_13">
      <!-- <img class="image_29" src="/img/ros-hmci/ee5680e0a521215326befe295219014f.png" /> -->
      <span class="pos_15">人机协同智能操作系统整体架构</span>

      <div class="Ovearc">
        <div class="section_22 pos_19"></div>
        <div class="section_24 pos_25"></div>
        <div class="section_25 pos_28"></div>
        <div class="section_27 pos_35"></div>

        <img class="image_30 pos_24" src="/img/ros-hmci/39ea174f887d877ea76ba31228a6315d.png" />
        <img class="image_30 pos_21" src="/img/ros-hmci/f21dd041539c46ca20943e97f003c425.png" />
        <img class="image_32 pos_29" src="/img/ros-hmci/58c00cc05cb3a7614393125e1a0bd7ba.png" />
        <img class="image_33 pos_32" src="/img/ros-hmci/f21dd041539c46ca20943e97f003c425.png" />
        <img class="image_30 pos_31" src="/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png" />
        <img class="image_30 pos_34" src="/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png" />
        <img class="image_30 pos_39" src="/img/ros-hmci/c283a7269004ff2edbff7619d4346ae1.png" />
        <img class="image_34 pos_37" src="/img/ros-hmci/3d0cdf6c59b3734bdcaa492f430697ae.png" />


        <div class="flex-col section_26 space-y-14 pos_33">
          <div class="flex-row items-center group_25 space-x-12">
            <img class="image" src="/img/ros-hmci/81de95962288a1aa874c40ee5f622d90.png" />
            <span class="font_5 text_47">人机协同智能操作系统</span>
          </div>
          <span class="text_48">
            人机协同智能操作系统是元操作系统（Meta
            OS），它运行于传统的计算机操作系统之上，实现智能无人平台的传感器、执行器、通信系统等硬件资源的虚化，管理智能无人系统的感知、规划、推理等自主行为，并对人机自主协同和持续自主学习提供支撑，研究人机协同智能操作系统架构层次性功能设计、协同机制以及接口标准，并持续迭代进行设计优化。
          </span>
        </div>
        <div class="flex-col items-center space-y-2 pos_91">
          <div class="flex-col justify-start items-center section_23 pos_jgt" @mouseenter="() => xtjg(1)"
            @mouseleave="unxtjg(1)">
            <a class="font_15 text_43">作业</a>
            <a class="font_15 text_43">环境模型</a>
          </div>
          <div id="jgt1" style="display: none">
            <img style="position: relative; left: -178px; top: 129px; transform: rotate(180deg)" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative; top: 5px; right: 470px" class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                作业环境模型中记录了机器和算法可理解的地图和障碍物、自身的位姿和速度等信息，以及用于识别是否关注目标的目标特征库。其关键技术包括高效的地图构建与更新技术、基于多传感器的自身位姿与速度估计、目标特征的检索/匹配/评估技术等。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col items-center space-y-2 pos_92">
          <div class="flex-col justify-start items-center text-wrapper_6 pos_jgt" @mouseenter="() => xtjg(2)"
            @mouseleave="unxtjg(2)">
            <a class="font_15 text_43">感知层</a>
          </div>
          <div id="jgt2" style="display: none">
            <img style="position: relative; left: -161px; top: 129px; transform: rotate(180deg)" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative;top: -33px;right: 453px;height: 154px;"
              class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                感知层不断综合各种传感器数据，基于各类感知算法，构建无人系统的作业环境模型，并确保作业环境模型总是被实时更新。其关键技术包手异构传感器融合框架以支撑多传感器数据融合技术，目标检测\识别\跟踪与障碍物检测技术、随机复杂高动态强鲁棒感知认知技术与算法等。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col items-center space-y-2 pos_93">
          <div class="flex-col items-center section_23 space-y-2 pos_jgt" @mouseenter="() => xtjg(3)"
            @mouseleave="unxtjg(3)">
            <a class="font_14">无人平台</a>
            <a class="font_14">硬件虚拟层</a>
          </div>
          <div id="jgt3" style="display: none">
            <img style="position: relative; left: -5px; top: 129px; z-index: 1" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative;top: -33px;right: -10px;z-index: 1;height: 142px;"
              class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                无人平台硬件虚拟层对无人平台各类硬件进行虚拟化，从而屏蔽硬件的物理差异，向上提供统一的虚拟模型和访问接口，为自主协同行为提供支撑。其关键技术包括各类无人平台的传感器及执行器的虚拟模型和访问接口设计、无线通信虚拟技术及其接口等。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col items-center space-y-2 pos_94">
          <div class="flex-col items-center section_23 space-y-2 pos_jgt" @mouseenter="() => xtjg(4)"
            @mouseleave="unxtjg(4)">
            <a class="font_13">虚拟空间</a>
            <a class="font_13">自主学习平台</a>
          </div>
          <div id="jgt4" style="display: none">
            <img style="position: relative; left: -5px; top: 129px" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative;top: -33px;right: -10px;height: 190px;"
              class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                虚拟空间自主学习平台运行在后台服务器集群中，用于构建虚拟空间，并实现并行可持续自主学习。其核心模块包括支持高性能、高精度、高保真、高动态的模拟仿真引擎，面向复杂环境和多型无人系统的数据生成引擎，虚拟空间自主学习引擎等。/匹配/评估技术等。针对实体空间的人机物全要素场景，研究构建全视景虚拟作业环境的智能方法和视景渲染技术。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col items-center space-y-2 pos_95">
          <div class="flex-col justify-start items-center text-wrapper_6 pos_jgt" @mouseenter="() => xtjg(5)"
            @mouseleave="unxtjg(5)">
            <a class="font_5 text_46">人机协同层</a>
          </div>
          <div id="jgt5" style="display: none">
            <img style="position: relative; top: 126px; right: 5px" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative;top: -33px;right: -10px;height: 230px;"
              class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                人机协同层支持人机自主协同，包括人机互理解模块和人机自主协同操作模块。人机互理解模块为人机自主协同提供界面、指令和基础算法，其关键技术包括作业指令设计、基于作业环境模型的人机交互以及协同目标识别等算法，协同作业规划模块实现人向无人系统的作业分配方法，作业预规划和动态作业规划等技。人机自主协同操控模块实现作业过程的人对无人系统的粗粒度控制，其关键技术包括自主协同操控通用框架、协同任务规划交互控制、协同路径规划与控制等技术。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col items-center space-y-2 pos_96">
          <div class="flex-col items-center section_29 space-y-2 pos_jgt" @mouseenter="() => xtjg(6)"
            @mouseleave="unxtjg(6)">
            <a class="font_13">规划</a>
            <a class="font_13">推理层</a>
          </div>
          <div id="jgt6" style="display: none">
            <img style="position: relative; top: 127px; right: 5px" class="image_31 pos_22"
              src="/img/ros-hmci/c5801299a6dacbe797ce18407faa827f.png" />
            <div style="position: relative; top: -33px; right: -10px;height:180px"
              class="flex-col justify-start items-start section_20">
              <span class="font_12 text_42">“</span>
              <span class="font_3 text_41 pos_17">
                规划与推理层基于感知信息和作业环境模型，按作业任务要求，开展高层次的规划和推理，将作业分解为可执行的状态机，并规划为无人平台可以执行的动作序列，由作业状态切换引擎和运动规划器组成。其关键技术包括作业状态机管理技术、基于状态机的推理技术、失效和异常处理技术等。运动规划器研究多目标融合决策与规划、长时多模协同混合控制等技术等。
              </span>
            </div>
          </div>
        </div>
        <div class="flex-col justify-start items-center text-wrapper_9 pos_97 small">
          <a class="text_97">执行器</a>
        </div>
        <div class="flex-col justify-start items-center text-wrapper_7 pos_98 small">
          <a class="text_98">运动规划</a>
        </div>
        <div class="flex-col justify-start items-center text-wrapper_7 pos_99 small">
          <a class="text_99">传感器</a>
        </div>
      </div>

    </div>



    <div class="pos_88">
      <span class="font_8 pos_44">系统典型应用场景</span>
      <span class="font_5 text_50 pos_45">

      </span>

      <div class="flex-row justify-center section_30 space-x-42" v-if="typical.length>=3">
        <div class="flex-col justify-start relative group_26">
          <div class="flex-col relative section_35">
            <span class="self-start font_16 text_51">{{typical[0].name}}</span>
            <span class="self-start font_17 text_54">{{typical[0].developers}}</span>
            <div class="flex-col items-start group_27">
              <span class="font_12 text_55">“</span>
              <span class="font_11 text_56 double-line"
                :title="typical[0].desc">
                {{typical[0].desc}}
              </span>
            </div>
            <div class="flex-row group_28 space-x-4">
              <a :href="typical[0].url" class="font_18">查看详情</a>
              <img class="image_38" src="/img/ros-hmci/5520ba673bc80210912d830a304adddd.png" />
            </div>
          </div>
          <img class="image_35 pos_49" src="/img/ros-hmci/cd2dbfa8c0212c06144c86f45ad70df8.png" />
          <div class="flex-col justify-start section_31 pos_46">
            <div class="section_32"></div>
          </div>
        </div>
        <div class="flex-col justify-start relative group_26">
          <div class="flex-col relative section_35">
            <span class="self-start font_16 text_52">{{ typical[1].name }}</span>
            <span class="self-start font_17 text_54">{{ typical[1].developers }}</span>
            <div class="flex-col items-start group_27">
              <span class="font_12 text_55">“</span>
              <span class="font_11 text_56 double-line"
                :title="typical[1].desc">
                {{typical[1].desc}}
              </span>
            </div>
            <div class="flex-row group_28 space-x-4">
              <a :href="typical[1].url" class="font_18">查看详情</a>
              <img class="image_38" src="/img/ros-hmci/5520ba673bc80210912d830a304adddd.png" />
            </div>
          </div>
          <img class="image_36 pos_50" src="/img/ros-hmci/7bd4e4622f6f137f0c4f348339b275c5.png" />
          <img class="image_37 pos_52" src="/img/ros-hmci/6f6bf64d608d346bd40464912be1df76.png" />
          <div class="flex-col justify-start section_33 pos_47">
            <div class="section_32"></div>
          </div>
        </div>
        <div class="flex-col justify-start relative group_26">
          <div class="flex-col relative section_35">
            <span class="self-start font_16 text_53">{{typical[2].name}}</span>
            <span class="self-start font_17 text_54">{{typical[2].developers}}</span>
            <div class="flex-col items-start group_27">
              <span class="font_12 text_55">“</span>
              <span class="font_11 text_56 double-line"
                :title="typical[2].desc">
                {{typical[2].desc}}
              </span>
            </div>
            <div class="flex-row group_28 space-x-4">
              <a :href="typical[2].url" class="font_18">查看详情</a>
              <img class="image_38" src="/img/ros-hmci/5520ba673bc80210912d830a304adddd.png" />
            </div>
          </div>
          <img class="image_36 pos_51" src="/img/ros-hmci/7bd4e4622f6f137f0c4f348339b275c5.png" />
          <img class="image_37 pos_53" src="/img/ros-hmci/6f6bf64d608d346bd40464912be1df76.png" />
          <div class="flex-col justify-start section_34 pos_48">
            <div class="shrink-0 section_32"></div>
          </div>
        </div>
      </div>

      <div>
        <div id="cbl1" class="pos_86">
          <div class="cbltext" v-show="showText[0]">
            <span class="text_96">意见反馈</span>
          </div>
          <a :href="links.home_advice"><img class="cblimg" :src="feedbackImgSrc" @mouseenter="handleMouseEnter(1)"
              @mouseleave="handleMouseLeave(1)" />
          </a>
        </div>

        <div id="cbl2" class="pos_85">
          <div class="cbltext" v-show="showText[1]" style="position: relative; top: 0px; left: -30px">
            <div style="display: flex; flex-direction: column; align-content: center">
              <img src="/img/ros-hmci/qrcode.png" />
              <span class="text_95">用户交流群</span>
            </div>
          </div>
          <img class="cblimg" style="position: absolute; right: -15px; top: -12px" :src="qrcodeImgSrc"
            @mouseenter="handleMouseEnter(2)" @mouseleave="handleMouseLeave(2)" />
        </div>
      </div>
    </div>

    <div class="flex-col justify-start section_36 pos_54">
      <div class="flex-col items-center section_37 space-y-50">
        <span class="font_8 text_58">现在开始！进入人机协同智能操作系统社区。</span>
        <div @click="getdownload('/os-system', 'osDownload')" class="flex-col justify-start items-center text-wrapper_10">
          <a class="font_10 text_59">立即下载</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavigationBar from "../components/NavigationBar.vue";
import axios from "axios";
import linksConfig from './links.js';

export default {
  name: "HomePage",
  data() {
    return {
      hoveredImage: -1,
      activeIndex: -1,
      statisticsData: {
        developer: "",
        organization: "",
      },
      links: linksConfig,
      showText: [false, false],
      feedbackImgSrc: "/img/ros-hmci/group7305.png",
      qrcodeImgSrc: "/img/ros-hmci/group7307.png",
      systemComponentTotal: 0,
      applicationSoftwareTotal: 0,
      typical: []
    };
  },
  components: {
    NavigationBar,
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
    this.fetchStatisticsData();
    this.fetchTypical();
    this.fetchTotal("ros-hmci-os", "systemComponentTotal");
    this.fetchTotal("ros-hmci-app", "applicationSoftwareTotal");
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  computed: {
    textVisible() {
      return {
        0: this.showText[0],
        1: this.showText[1]
      };
    }
  },
  methods: {
    async fetchStatisticsData() {
      try {
        //主页的项目协同后两项获取地址
        const response = await axios.get(this.links.home_statisticsData);
        /*
        json格式如下
        {
          "developer": "500+",
          "organization": "100+"
        }
        */
        this.statisticsData.developer = response.data.developer;
        this.statisticsData.organization = response.data.organization;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    async fetchTypical() {
      try {
        const response = await fetch(
          `/fanshuai/ROS_hmci_Typical_app/raw/branch/master/typical.json`
        )
        const data = await response.json();
        this.typical = data;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    handleMouseEnter(id) {
      if (id === 1) {
        this.showText[0] = true;
        this.feedbackImgSrc = "/img/ros-hmci/mbz668.png";
      } else if (id === 2) {
        this.showText[1] = true;
        this.qrcodeImgSrc = "/img/ros-hmci/mbz669.png";
      }
    },

    handleMouseLeave(id) {
      if (id === 1) {
        this.showText[0] = false;
        this.feedbackImgSrc = "/img/ros-hmci/group7305.png";
      } else if (id === 2) {
        this.showText[1] = false;
        this.qrcodeImgSrc = "/img/ros-hmci/group7307.png";
      }
    },

    async fetchTotal(topic, targetDataField) {
      try {
        const response = await axios.get(
          `/explore/repos/search?topic=${encodeURIComponent(topic)}`
        );
        this[targetDataField] = response.data.Data.Total;
      } catch (error) {
        console.error("接口错误", error);
      }
    },

    getdownload(path, id) {
      var path1 = path;
      var Id = id;
      this.$router.push({
        path: path1,
        query: {
          id: Id,
        },
      });
    },

    mouseEnter(index) {
      this.hoveredImage = index;
      this.activeIndex = index;
      const tabbar = document.querySelector(".tabBar");
      const tabIds = ["#tab1", "#tab2", "#tab3", "#tab4"];

      for (let i = 0; i < tabIds.length; i++) {
        if (i === index) {
          document.querySelector(tabIds[i]).style.display = "";
        } else {
          document.querySelector(tabIds[i]).style.display = "none";
        }
      } const positions = [50, 350, 640, 974];
      tabbar.style.left = positions[index] + "px";
    },

    xtjg(index) {
      const xgtIds = ["#jgt1", "#jgt2", "#jgt3", "#jgt4", "#jgt5", "#jgt6"];

      for (let i = 0; i < xgtIds.length; i++) {
        if (i === index - 1) {
          document.querySelector(xgtIds[i]).style.display = "";
        } else {
          document.querySelector(xgtIds[i]).style.display = "none";
        }
      }
    },
    unxtjg(index) {
      const xgtIds = ["#jgt1", "#jgt2", "#jgt3", "#jgt4", "#jgt5", "#jgt6"];

      for (let i = 0; i < xgtIds.length; i++) {
        if (i === index - 1) {
          document.querySelector(xgtIds[i]).style.display = "none";
        } else {
          document.querySelector(xgtIds[i]).style.display = "none";
        }
      }
    },
  },
};
</script>

<style scoped>
.flexCont{
  display: flex;
  align-items: center;
}
.section_6_cont{
  width: 1200px;
  height: 100%;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.homeHeadCont{
  color:#363737;
  font-size:40px;
  font-weight: bold;
  line-height: normal;
  letter-spacing: 4px;
  position: relative;
  top: -50px;
}
.homeBox1, .homeBox3, .gotoDownload{
  font-weight: normal;
}
.homeBox1{
  color:#ffffff;
  font-size:21px;
  background-image:linear-gradient(133.55deg,#edbd4a 0%,#19bf77 100%);
  border-radius:4px;
  padding: 8px 12px;
  margin-left: 12px;
}
.homeBox2{
  background-image:-webkit-linear-gradient(left,#FAC171,#02E588);
  -webkit-background-clip:text;
  -webkit-text-fill-color:transparent;
}
.homeBox3{
  color:#ffffff;
  font-size:20px;
  background-image:linear-gradient(150.44deg,#19bf77 0%,rgba(26, 209, 195, 0) 100%);
  border-radius: 6px;
  padding: 10px 150px 10px 25px;
  margin-top: 20px;
}
.gotoDownload{
  background-color: #19bf77;
  padding: 15px 30px;
  margin-top: 45px;
  display: inline-block;
  border-radius: 4px;
}
.banner_right{
  position: relative;
  height: 100%;
  flex: 1;
  right: -5vw;
  display: inline-flex;
  justify-content: center;
  align-items: center;
}
.homepage1{
  width: 230px;
  z-index: 2;
  animation: scaleDraw 4s ease-in-out infinite;
  -webkit-animation: scaleDraw 4s ease-in-out infinite;
}
@keyframes scaleDraw {
  0% {transform:scale(1);}
  25% {transform:scale(1.3);}
  50% {transform:scale(1);}
  75% {transform:scale(1.3);}
}
.homepage2{
  width: 100px;
  z-index: 2;
  position: absolute;
  top: 7vh;
  left: 25px;
  /* animation: bounce 3s ease-in-out infinite;
  -webkit-animation: bounce 3s ease-in-out infinite; */
}
.homepage3{
  width: 120px;
  z-index: 2;
  position: absolute;
  bottom: 10vh;
  right: 0;
  /* animation: bounce 2s ease-in-out infinite;
  -webkit-animation: bounce 2s ease-in-out infinite; */
}
@keyframes bounce {

  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-20px);
  }
}
.homepage4{
  width: 716px;
  z-index: 0;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
}
</style>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.double-line {
  word-break: break-all;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}


@keyframes fadeIn {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}


@keyframes changeHeight {

  0%,
  100% {
    transform: scaleY(1);
  }

  25% {
    transform: scaleY(0.8);
  }

  50% {
    transform: scaleY(1.2);
  }

  75% {
    transform: scaleY(0.9);
  }
}

.sec_1,
.sec_2,
.sec_3,
.sec_4 {
  animation: changeHeight 5s infinite;
  transform-origin: bottom;
}

.sec_1 {
  position: absolute;
  top: 220px;
  right: 963px;
  height: 24px;
}

.sec_2 {
  animation-delay: 1s;
  position: absolute;
  top: 208px;
  right: 950px;
  height: 42px;
}

.sec_3 {
  animation-delay: 1.5s;
  position: absolute;
  top: 222px;
  right: 936px;
  height: 34px;
}

.sec_4 {
  animation-delay: 2s;
  position: absolute;
  top: 207px;
  right: 923px;
  height: 56px;
}



@keyframes bubble {

  0%,
  100% {
    transform: scale(1);
  }

  50% {
    transform: scale(1.1);
  }
}

.bubble-animation {
  animation: bubble 0.5s linear;
}

.banner_a {
  position: relative;
}

.text_31 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 18px;
  text-align: center;
}


.banner_bg {
  height: 186px;
  width: 322px;
  position: absolute;
  top: 69px;
  right: 492px;
  animation: fadeinout 1.5s ease-in-out infinite;
}

@keyframes fadeinout {
  0% {
    opacity: 0;
  }

  50% {
    opacity: 1;
  }

  100% {
    opacity: 0;
  }
}

.bounce {
  animation: bounce 1.5s infinite;
}

@keyframes bounce {

  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-10px);
  }
}

.banner_ros {
  position: absolute;
  top: 90px;
  right: 581px;
  width: 156.5px;
  height: 85px;
}

.Ovearc {
  position: relative;
  top: 133px;
}

.page {
  /* background-color: #0b0e1d; */
  width: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  height: 4116px;
  margin-bottom: -80px;
}

.section {
  padding: 17px 80px 14px 123px;
  filter: drop-shadow(0px 2px 3px #0137931f);
  background-color: #282828;
}

.image_4 {
  width: 63px;
  height: 16px;
}

.image_41 {
  width: 20px;
  height: 20px;
  margin-right: 5px;
}

.image_42 {
  width: 50px;
  height: 50px;
}

.section_3 {
  margin-left: 14px;
  background-color: #ffffff;
  width: 2px;
  height: 20.5px;
}

.text {
  margin-left: 18px;
  color: #ffffff;
  font-size: 22px;
  font-family: AlimamaShuHeiTi;
  line-height: 20.5px;
  letter-spacing: 1px;
}

.group_3 {
  margin-left: 48px;
  height: 25.5px;
}

.space-x-39>*:not(:first-child) {
  margin-left: 39px;
}

.image_5 {
  width: 29px;
  height: 26px;
}

.font_1 {
  font-size: 16px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #ffffff;
}

.group {
  width: 104px;
}

.section_2 {
  border-radius: 15px;
  opacity: 0.7;
  height: 30px;
  border: solid 1px #ffffff;
}

.text_12 {
  font-size: 15px;
  line-height: 13.5px;
  opacity: 0.63;
}

.pos_2 {
  position: absolute;
  right: 23.5px;
  top: 50%;
  transform: translateY(-50%);
}

.image_7 {
  opacity: 0.63;
  width: 14px;
  height: 14px;
}

.pos {
  position: absolute;
  left: 25px;
  top: 50%;
  transform: translateY(-50%);
}

.group_2 {
  margin-left: 20px;
  width: 28.5px;
  height: 25px;
}

.image_6 {
  margin-top: 6px;
  opacity: 0.7;
  width: 21px;
  height: 19px;
}

.text-wrapper {
  margin-top: -25px;
  padding: 4px 0;
  background-color: #d92b2f;
  border-radius: 50%;
  width: 12px;
}

.text_2 {
  color: #ffffff;
  font-size: 8px;
  font-family: Alibaba PuHuiTi;
  line-height: 5.5px;
}

.image_3 {
  margin-left: 12px;
  opacity: 0.7;
  width: 20px;
  height: 20px;
}

.section_4 {
  margin-left: 20px;
  background-color: #ffffff7a;
  width: 1px;
  height: 13px;
}

.image {
  width: 30px;
  height: 30px;
}

.image_2 {
  margin-left: 20px;
}

.font_2 {
  font-size: 14px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #ffffff;
}

.text_3 {
  opacity: 0.9;
}

.text_11 {
  font-size: 15px;
  line-height: 12px;
}

.text_10 {
  font-size: 15px;
  line-height: 12px;
}

.text_9 {
  font-size: 15px;
  line-height: 12px;
}

.text_8 {
  font-size: 15px;
  line-height: 12px;
}

.text_7 {
  font-size: 15px;
  line-height: 12px;
}

.text_72 {
  font-family: PingFang SC;
  color: #ffffff;
  font-size: 18px;
  text-align: center;
}

.text_6 {
  font-size: 15px;
  line-height: 12px;
}

.text_5 {
  font-size: 15px;
  line-height: 12px;
}

.text_4 {
  font-size: 15px;
  line-height: 12px;
}

.text_13 {
  margin-left: 10px;
}

.image_9 {
  margin-left: 8px;
  border-radius: 1px;
  width: 11px;
  height: 6px;
}

.image_8 {
  margin-left: 20px;
  opacity: 0.7;
  width: 20px;
  height: 19px;
}

/* .group_4 {
  padding-bottom: 663px;
  overflow-x: hidden;
} */

.section_6 {
  position: relative;
  top: 0px;
  height: 600px;
  background-image: url("/img/ros-hmci/homepageBanner.png");
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
}

.section_5 {
  position: relative;
  top: 27px;
  font-size: 14px;
  padding: 10px 18.75%;
  background-color: rgba(255, 255, 255, 0.21);
}

.space-x-14>*:not(:first-child) {
  margin-left: 14px;
}

.image_10 {
  width: 33px;
  height: 21px;
}

.font_3 {
  font-size: 14px;
  font-family: Alibaba PuHuiTi;
}

.text_14 {
  font-family: Alibaba PuHuiTi;
  color: white;
  font-size: 14px;
}

.group_5 {
  position: relative;
  left: 18.75%;
  top: 110px;
}

.font_4 {
  font-size: 38px;
  font-family: Alibaba PuHuiTi;
  line-height: 36px;
  color:#363737;
}

.text_15 {
  line-height: 35.5px;
}

.group_6 {
  margin-top: 16px;
}

.image_11 {
  width: 146px;
  height: 35px;
}

.font_5 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 14.5px;
  color: #ffffff;
}

.text_16 {
  font-family: Alibaba PuHuiTi;
  color: #ffffff;
  font-size: 20px;
}

.text-wrapper_2 {
  padding: 14px 0;
  background-image: linear-gradient(156.43deg, #35cf67 0%, #179d8c 100%);
  border-radius: 4px;
  width: 131px;
  border: solid 1px #00000000;
  position: relative;
  left: 18.75%;
  top: 185px;
}

.text-wrapper_2:hover {
  background-image: linear-gradient(203.57deg, #35cf67 0%, #179d8c 100%);
}


.image_12 {
  margin-left: 18.75%;
  margin-top: 68px;
  width: 84px;
  height: 2px;
}

.section_7 {
  margin-top: 80px;
  padding: 32px 18.75% 22px;
  background-color: rgba(255, 255, 255, 1);
  /* backdrop-filter: blur(60px); */
}

.space-x-24>*:not(:first-child) {
  margin-left: 24px;
}

.image_13 {
  margin-top: 2px;
  width: 45px;
  height: 39px;
}

.font_6 {
  font-size: 28px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #ffffff;
}

.text_17 {
  line-height: 20px;
}

.text_21 {
  margin-right: 4px;
  font-size: 15px;
  line-height: 14px;
}

.group_7 {
  width: 896px;
}

.space-x-8>*:not(:first-child) {
  margin-left: 8px;
}

.section_8 {
  background-color: #caced8;
  width: 1px;
  height: 51px;
}

.group_8 {
  width: 321.5px;
}

.space-x-22>*:not(:first-child) {
  margin-left: 22px;
}

.image_14 {
  margin-left: 98px;
  width: 39px;
  height: 36px;
}

.text_18 {
  line-height: 20px;
}

.text_22 {
  font-size: 15px;
  line-height: 14px;
}

.group_9 {
  width: 565px;
}

.view {
  margin-bottom: 2px;
}

.group_10 {
  margin-left: 14px;
}

.group_11 {
  width: 310px;
}

.image_15 {
  margin-left: 82px;
  margin-top: 2px;
  width: 41px;
  height: 39px;
}

.text_19 {
  line-height: 20px;
}

.text_23 {
  font-size: 15px;
  line-height: 14px;
}

.section_9 {
  margin-bottom: 2px;
  background-color: #caced880;
  width: 1px;
  height: 51px;
}

.group_12 {
  margin-left: 104px;
  margin-top: 2px;
}

.image_16 {
  margin-top: 2px;
  width: 49px;
  height: 41px;
}

.text_20 {
  line-height: 20px;
}

.text_24 {
  margin-right: 28px;
  font-size: 15px;
  line-height: 14px;
}

.section_30 {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  bottom: 125px;
}

.space-x-42>*:not(:first-child) {
  margin-left: 42px;
}

.group_26 {
  padding-top: 28px;
  width: 372.5px;
}

.group_26:hover {
  box-shadow: 5px 20px 40px rgba(0, 0, 0, 0.03922);
  transform: translateY(-8px);
  transition: transform 0.3s ease-in-out;
}

.section_35 {
  padding: 35px 20px 26px;
  background-color: #ffffff;
  border-radius: 8px;

  border: solid 1px #ebf7f6;
  height: 278px;
}

.font_16 {
  font-size: 22px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #181818;
}

.text_51 {
  line-height: 21px;
}

.font_17 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 14.5px;
  color: #a4a9b5;
}

.text_54 {
  margin-top: 20px;
}

.group_27 {
  margin-top: 40px;
}

.font_12 {
  font-size: 40px;
  font-family: Alibaba PuHuiTi;
  line-height: 10px;
  color: #24a19b;
}

.text_55 {
  margin-left: -23px;
}

.info-div {
  margin-top: 17px;
  display: inline-flex;
  flex-direction: row;
  align-items: center;
  padding-left: 23.45px;
  width: 419px;
  height: 42px;
  background-image: linear-gradient(90deg, #1cb08f 0%, rgba(3, 13, 35, 0.46) 100%);
  border-radius: 4px;
}

.font_11 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #4d4d4d;
}

.text_56 {
  margin-top: 12px;
}

.text_57 {
  margin-top: 2px;
}

.group_28 {
  position: absolute;
  bottom: 26px;
}

.space-x-4>*:not(:first-child) {
  margin-left: 9px;
}

.font_18 {
  font-size: 17px;
  font-family: Alibaba PuHuiTi;
  line-height: 14.5px;
  color: #24a19b;
}

.image_38 {
  position: absolute;
  left: 63px;
  bottom: -1px;
  width: 25px;
  height: 13px;
}

.image_35 {
  width: 115px;
  height: 146px;
}

.pos_49 {
  position: absolute;
  right: 0;
  top: 5px;
}

.section_31 {
  background-color: #ffffff;
  border-radius: 50%;
  background-image: url("/img/ros-hmci/9232cbc26c8e54585239620bfa45450d.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 100px;
}

.pos_46 {
  position: absolute;
  right: 40.5px;
  top: 0;
}

.section_32 {
  border-radius: 50%;
  height: 100px;
  border: solid 1px #e5b372;
}

.text_52 {
  line-height: 20.5px;
}

.image_36 {
  width: 115px;
  height: 112px;
}

.pos_50 {
  position: absolute;
  right: 0;
  top: 5.5px;
}

.image_37 {
  width: 59px;
  height: 57px;
}

.pos_52 {
  position: absolute;
  right: 14px;
  top: 94px;
}

.section_33 {
  background-image: url("/img/ros-hmci/abbaa2027cb92558c1fa576750dfe8da.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 100px;
  height: 100px;
}

.pos_47 {
  position: absolute;
  right: 40.5px;
  top: 0;
}

.text_53 {
  line-height: 20px;
}

.pos_51 {
  position: absolute;
  right: 0;
  top: 5px;
}

.pos_53 {
  position: absolute;
  right: 14px;
  top: 93.5px;
}

.section_34 {
  background-image: url("/img/ros-hmci/b808f935653d2ec8945a3e16f4af254c.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 100px;
}

.pos_48 {
  position: absolute;
  right: 40.5px;
  top: 0;
}

.section_10 {
  height: 1850px;
  background-image: url("/img/ros-hmci/mbz667.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
}

.pos_3 {
  position: absolute;
  left: 0;
  right: -1px;
  top: 600px;
}

.font_8 {
  font-size: 36px;
  font-family: Alibaba PuHuiTi;
  color: #181818;
}

.font_7 {
  font-size: 50px;
  font-family: PingFang SC;
  color: #24a19b;
}

.text_25 {
  opacity: 0.11;
}

.pos_4 {
  position: absolute;
  left: 50%;
  top: 60px;
  transform: translateX(-50%);
}

.cblimg {
  width: 58px;
  height: 58px;
}

.cbltext {
  height: auto;
  width: auto;
  flex-direction: row;
  padding: 6.5px 10px;
  background-color: #ffffff;
  border-radius: 4px;
}

.pos_85 {
  width: 175px;
  height: 147px;
  align-items: center;
  justify-content: center;
  display: inline-flex;
  flex-direction: row;
  position: absolute;
  right: 79px;
  bottom: 15px;
}

.pos_86 {
  align-items: center;
  justify-content: center;
  display: inline-flex;
  flex-direction: row;
  position: absolute;
  right: 65px;
  bottom: 182px;
}

.pos_7 {
  position: absolute;
  top: 1150px;
  left: 50%;
  width: 1118px;
  transform: translateX(-50%);
  justify-content: space-between;
}

.image_24 {
  margin-top: 80px;
  width: 469px;
  height: 147px;
}

.space-x-110>*:not(:first-child) {
  margin-left: 97px;
}

.image_21 {
  width: 75px;
  height: 77px;
}

.font_10 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 18px;
  text-align: center;
}

.text_30 {
  margin-left: -32px;
  color: #24a19b;
}

.section_14 {
  background-image: linear-gradient(102.9deg, #1ad1c3 6.2%, #02e588 51.1%);
  border-radius: 20px;
  width: 41px;
  height: 5px;
  border: solid 1px #00000000;
}

.space-x-92>*:not(:first-child) {
  margin-left: 92px;
}

.image_22 {
  width: 72px;
  height: 79px;
}

.image_19 {
  width: 76px;
  height: 85px;
}

.text_32 {
  margin-left: -14px;
}

.space-y-10>*:not(:first-child) {
  margin-top: 10px;
}

.space-y-18>*:not(:first-child) {
  margin-top: 18px;
}

.image_23 {
  width: 79px;
  height: 75px;
}

.text_33 {
  margin-left: -48px;
  line-height: 17px;
}

.group_14 {
  align-items: center;
}

.group_15 {
  align-items: center;
  margin-left: 59px;
}

.group_17 {
  align-items: center;
  margin-left: 60px;
}

.group_19 {
  align-items: center;
  margin-left: 51px;
}

.space-y-12>*:not(:first-child) {
  margin-top: 12px;
}

.image_20 {
  margin-left: -6px;
  width: 70px;
  height: 85px;
}

.text_34 {
  margin-left: -68px;
}

.section_12 {
  background-color: #edf5fc;
  height: 309px;
}

.pos_6 {
  position: absolute;
  left: 0;
  right: 1279px;
  top: 947px;
}

.section_19 {
  padding: 260px 0 762px;
  background-image: url("/img/ros-hmci/架构底部.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
}

.pos_13 {
  position: relative;
  left: 0;
  right: 0;
  top: 1500px;
  height: 1153px;
}

.image_29 {
  width: 490px;
  height: 34px;
}

.pos_15 {
  font-size: 36px;
  color: #ffffff;
  position: absolute;
  left: 50%;
  top: 265px;
  transform: translateX(-50%);
}

.section_22 {
  opacity: 0.3;
  border-radius: 50%;
  height: 857px;
  border: solid 1px #caced8;
}

.pos_19 {
  width: 857px;
  position: absolute;
  top: -20px;
  left: 50%;
  transform: translateX(-50%);
}

.pos_44 {
  position: absolute;
  left: 50%;
  top: 65px;
  transform: translateX(-50%);
}

.text_50 {
  color: #757c8d;
  line-height: 15px;
}

.pos_45 {
  position: absolute;
  left: 50%;
  top: 128px;
  transform: translateX(-50%);
}

.section_27 {
  border-radius: 50%;
  opacity: 1;
  height: 308px;
  border: solid 1px #caced8;
}

.pos_35 {
  width: 308px;
  position: absolute;
  top: 245px;
  left: 50%;
  transform: translateX(-50%);
}

.section_25 {
  opacity: 0.7;
  border-radius: 50%;
  height: 482px;
  border: solid 1px #caced8;
}

.pos_28 {
  width: 482px;
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  top: 155px;
}

.section_24 {
  opacity: 0.55;
  border-radius: 50%;
  height: 664px;
  border: solid 1px #caced8;
}

.pos_25 {
  width: 664px;
  position: absolute;
  top: 63px;
  left: 50%;
  transform: translateX(-50%);
}

.section_21 {
  padding: 38px 0;
  filter: drop-shadow(0px 3px 7.5px #1078a769);
  background-image: url("/img/ros-hmci/981ec94220a67dd1b9a125ed52371d75.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 116px;
}

.space-y-4>*:not(:first-child) {
  margin-top: 4px;
}

.pos_18 {
  position: absolute;
  left: 821.5px;
  top: 2137.5px;
}

.font_13 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color: #4d4d4d;
}

.image_30 {
  width: 17px;
  height: 17px;
}

.pos_24 {
  position: absolute;
  right: -43px;
  top: 149px;
}

.pos_21 {
  position: absolute;
  left: -232px;
  bottom: -50px;
}

.image_32 {
  opacity: 0.57;
  width: 17px;
  height: 17px;
}

.pos_29 {
  position: absolute;
  left: 417px;
  bottom: -465px;
}

.image_33 {
  opacity: 0.59;
  width: 17px;
  height: 17px;
}

.pos_32 {
  position: absolute;
  right: -290px;
  bottom: -227px;
}

.pos_31 {
  position: absolute;
  left: 1088px;
  bottom: 438px;
}

.pos_34 {
  position: absolute;
  right: 414px;
  bottom: -350px;
}

.pos_39 {
  position: absolute;
  right: 1078px;
  bottom: 682px;
}

.image_34 {
  opacity: 0.46;
  width: 17px;
  height: 17px;
}

.pos_37 {
  position: absolute;
  left: -303px;
  bottom: -250px;
}

.section_29 {
  padding: 32px 0 26px;
  background-image:linear-gradient(189.96deg,#bdff91 0%,#91ffd7 37.65%,#87eee9 100%);
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.64),0px 3px 20px #85f2eb inset;
  border-radius: 50%;
  width: 99px;
}

.section_28 {
  padding: 38px 0 36px;
  background-image: linear-gradient(190deg, #1b283b 15.1%, #233a48 92.5%);
  box-shadow: 0px 3px 15px #1078a769;
  border-radius: 50%;
  width: 116px;
}

.pos_42 {
  position: absolute;
  right: 621px;
  bottom: 1434.5px;
}

.text-wrapper_8 {
  padding: 42px 0;
  background-image: linear-gradient(190deg, #1b283b 15.1%, #233a48 92.5%);
  box-shadow: 0px 3px 15px #1078a769;
  border-radius: 50%;
  width: 100px;
}

.pos_27 {
  position: absolute;
  right: 785px;
  top: 2346.5px;
}

.font_15 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 14.5px;
  color:#4d4d4d;
}

.text_45 {
  line-height: 15px;
}

.text-wrapper_9 {
  padding: 26.5px 0;
  background-image:linear-gradient(189.96deg,#bdff91 0%,#91ffd7 37.65%,#87eee9 100%);
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.64),0px 3px 20px #85f2eb inset;
  border-radius: 50%;
  width: 74px;
}

.pos_40 {
  position: absolute;
  left: 627x;
  top: 2588.5px;
}

.section_23 {
  width: 156px;
  padding: 58px 0;
  background-image:linear-gradient(180deg,#ffffff 0%,#cdf6ae 100%);
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.35),0px 3px 20px #ffffff inset;
  border-radius: 50%;
}

.space-y-2>*:not(:first-child) {
  margin-top: 2px;
}

.pos_38 {
  position: absolute;
  left: 461.5px;
  top: 2557.5px;
}

.pos_41 {
  position: absolute;
  right: 490px;
  top: 2621.5px;
}

.pos_20 {
  position: absolute;
  right: 677px;
  top: 2169.5px;
}

@keyframes scaleDraw {
  0% {
    transform: scale(1);
  }

  25% {
    transform: scale(0.9);
  }

  50% {
    transform: scale(1);
  }

  75% {
    transform: scale(0.9);
  }
}

.small {
  -webkit-animation-name: scaleDraw;
  /*关键帧名称*/
  -webkit-animation-timing-function: ease-in-out;
  /*动画的速度曲线*/
  -webkit-animation-iteration-count: infinite;
  /*动画播放的次数*/
  -webkit-animation-duration: 5s;
  /*动画所花费的时间*/
}

.pos_jgt {
  position: absolute;
  right: 300px;
  top: 60px;
  -webkit-animation-name: scaleDraw;
  /*关键帧名称*/
  -webkit-animation-timing-function: ease-in-out;
  /*动画的速度曲线*/
  -webkit-animation-iteration-count: infinite;
  /*动画播放的次数*/
  -webkit-animation-duration: 5s;
  /*动画所花费的时间*/
}

.pos_jgt:hover {
  box-shadow: 0px 3px 15px rgba(16, 120, 167, 0.97);
}

.font_14 {
  font-size: 16px;
  font-family: Alibaba PuHuiTi;
  line-height: 19px;
  color:#4d4d4d;
}

.text-wrapper_6 {
  padding: 58px 0;
  background-image:linear-gradient(189.96deg,#bdff91 0%,#91ffd7 37.65%,#87eee9 100%);
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.64),0px 3px 20px #85f2eb inset;
  border-radius: 50%;
  width: 135px;
}

.pos_30 {
  position: absolute;
  right: 489px;
  top: 2374.5px;
}

.text_46 {
  color: #4d4d4d;
  line-height: 15px;
}

.pos_23 {
  position: absolute;
  left: 530.5px;
  top: 2258.5px;
}

.text_43 {
  line-height: 19px;
}

.section_26 {
  padding: 63px 90px 82px 98px;
  background-image: url("/img/ros-hmci/b013c67df9284efc84b92a51a5cb6d13.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 492px;
}

.space-y-14>*:not(:first-child) {
  margin-top: 14px;
}

.pos_33 {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  top: 245px;
  margin: 0 auto;
}

.group_25 {
  padding: 0 4px;
}

.space-x-12>*:not(:first-child) {
  margin-left: 12px;
}

.text_47 {
  color: #1ad1c3;
  line-height: 15px;
}

.text_48 {
  color: #4d4d4d;
  font-size: 13px;
  font-family: Alibaba PuHuiTi;
  font-weight: 400;
  line-height: 16px;
}

.text-wrapper_7 {
  padding: 36px 0;
  background-image:linear-gradient(180deg,#ffffff 0%,#cdf6ae 100%);
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.35),0px 3px 20px #ffffff inset;
  border-radius: 50%;
  width: 93px;
}

.pos_26 {
  position: absolute;
  left: 734.5px;
  top: 2380.5px;
}

.text_44 {
  color: #f95f9f;
  line-height: 15px;
  text-shadow: 0px 3px 15px #19bf774d;
}

.group_20 {
  padding: 43px 0;
  width: 1200px;
}

.pos_9 {
  position: absolute;
  left: 50%;
  top: 1330px;
  transform: translateX(-50%);
}

.section_17 {
  padding-top: 45px;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0px 3px 12px #0000000a;
  width: 648px;
}

.group_21 {
  padding-bottom: 20px;
}

.text_36 {
  margin-left: 210px;
  width: 401px;
  line-height: 27px;
}

.font_9 {
  font-size: 50px;
  font-family: Alibaba PuHuiTi;
  line-height: 13px;
}

.text_35 {
  position: absolute;
  left: 180px;
  top: 64px;
  color: #caced8;
}

.group_22 {
  padding-bottom: 44px;
}

.text-wrapper_3 {
  margin-left: 209px;
  padding: 14px 0;
  background-color: rgba(246, 174, 80, 0.35);
  border-radius: 4px;
  width: 101px;
}

.text_37 {
  color: #e59834;
  font-size: 15px;
  line-height: 14px;
}

.text-wrapper_5 {
  margin-left: 208px;
  margin-top: 72px;
  padding: 12px 0;
  background-image: linear-gradient(156.43deg, #35cf67 0%, #179d8c 100%);
  border-radius: 4px;
  width: 131px;
  border: solid 1px #00000000;
}

.text-wrapper_5:hover {
  background-image: linear-gradient(169.45deg, #35cf67 0%, #3dc1bb 100%);
}

.text_39 {
  line-height: 15px;
}

.group_23 {
  width: 320px;
}

.pos_11 {
  position: absolute;
  right: 0;
  top: 0;
}

.image_26 {
  opacity: 1.15;
  width: 269px;
  height: 219px;
}

.text-wrapper_4 {
  padding: 14px 0;
  background-color: #f6ae5059;
  border-radius: 4px;
  width: 130px;
}

.pos_12 {
  position: absolute;
  left: 0;
  top: 0;
}

.text_38 {
  color: #e59834;
  font-size: 15px;
  font-family: Alibaba PuHuiTi;
  line-height: 13.5px;
}

.section_15 {
  padding: 20px 0;
  background-image: linear-gradient(213.8deg, #35cf6700 3.2%, #179d8c00 74.7%);
  border-radius: 12px;
  box-shadow: 0px 3px 12px #0000000a;
  backdrop-filter: blur(17px);
  width: 720px;
  border: solid 1px #ffffff;
}

.pos_10 {
  position: absolute;
  left: 0;
  top: 0;
}

.tab_sec1 {
  margin-left: 18px;
  background-image: url("/img/ros-hmci/视频图1.png");
  background-size: 100% 100%;
  width: 679px;
  height: 430px;
}

.tab_sec2 {
  margin-left: 18px;
  background-image: url("/img/ros-hmci/视频图2.png");
  background-size: 100% 100%;
  width: 679px;
  height: 430px;
}

.tab_sec3 {
  margin-left: 18px;
  background-image: url("/img/ros-hmci/视频图3.png");
  background-size: 100% 100%;
  width: 679px;
  height: 430px;
}

.tab_sec4 {
  margin-left: 18px;
  background-image: url("/img/ros-hmci/视频图4.png");
  background-size: 100% 100%;
  width: 679px;
  height: 430px;
}

.space-y-134>*:not(:first-child) {
  margin-top: 134px;
}

.image_25 {
  margin-right: 270px;
  width: 57px;
  height: 57px;
}

.group_24 {
  padding-bottom: 10px;
}

.space-x-20>*:not(:first-child) {
  margin-left: 20px;
}

.image_27 {
  width: 10px;
  height: 12px;
}

.divider {
  background-color: #ffffff80;
  border-radius: 20px;
}

.section_18 {
  background-color: #ffffff;
  border-radius: 20px;
  width: 284px;
  height: 4px;
}

.image_28 {
  opacity: 0.05;
  filter: blur(44px);
  border-radius: 50%;
  width: 419px;
  height: 336px;
}

.pos_36 {
  position: absolute;
  right: -25px;
  bottom: 1226.5px;
}

.pos_14 {
  position: absolute;
  right: 190.5px;
  bottom: 1792.5px;
}

.section_20 {
  padding: 20px 0 144px;
  background-image:linear-gradient(180deg,#ffffff 0%,#fdffe2 100%);
  border-radius:8px;
  box-shadow:0px 3px 15px rgba(0, 136, 145, 0.59),0px 3px 5px #ffffff inset;
  width: 292px;
}

.pos_16 {
  position: absolute;
  right: 359px;
  top: 2078.5px;
}

.text_42 {
  margin-left: 0px;
}

.text_41 {
  color: #a4a9b5;
  line-height: 17px;
  text-indent: 25.5px;
}

.pos_17 {
  position: absolute;
  left: 15.5px;
  right: 11.5px;
  top: 50%;
  transform: translateY(-50%);
}

.image_31 {
  border-radius: 2px;
  width: 16px;
  height: 22px;
}

.pos_22 {
  position: absolute;
  right: 647px;
  bottom: 1866px;
}

.section_11 {
  padding: 48px 38px 20px;
  background-image: url("/img/ros-hmci/简介底部.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
  width: 1200px;
}

.space-x-40>*:not(:first-child) {
  margin-left: 40px;
}

.pos_5 {
  position: absolute;
  left: 50%;
  top: 182px;
  transform: translateX(-50%);
}

.group_13 {
  margin-top: 26px;
  word-spacing: -1em;
  height: 160.5px;
}

.text_28 {
  position: relative;
  top: 18px;
  left: 8px;
  color: #24a19b;
  vertical-align: sub;
  word-spacing: normal;
}

.text_27 {
  color: #757c8d;
  font-size: 22px;
  font-family: Alibaba PuHuiTi;
  line-height: 36px;
  word-spacing: normal;
}

.image_17 {
  width: 269px;
  height: 248px;
}

.text_49 {
  line-height: 46.5px;
}

.pos_43 {
  position: absolute;
  left: 50%;
  top: 36px;
  transform: translateX(-50%);
}

.tabBar {
  position: absolute;
  left: 144px;
  top: 140px;
  background-image: linear-gradient(102.9deg, #1ad1c3 6.2%, #02e588 51.1%);
  border-radius: 20px;
  width: 41px;
  height: 5px;
  border: solid 1px #00000000;
  z-index: 999;
  transition: all 0.2s;
}

.pos_8 {
  position: absolute;
  left: 407px;
  top: 1152.5px;
  transition: all 0.3s;
}

.image_18 {
  width: 101px;
  height: 101px;
}

.text_29 {
  margin-left: -46px;
  line-height: 17px;
}

.section_36 {
  background-image: url("/img/ros-hmci/开始底部.png");
  background-size: 100% 100%;
  background-repeat: no-repeat;
}

.pos_54 {
  position: absolute;
  left: 0;
  right: 0;
  top: 3860px;
}

.pos_56 {
  position: absolute;
  top: 2500px;
  left: 603px;
}

.text_95 {
  text-align: center;
  padding-top: 5px;
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 16px;
}

.text_96 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 16px;
}

.text_97 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 16px;
  text-align: center;
}

.text_98 {
  text-shadow: 0px 3px 15px rgba(25, 191, 119, 0.3);
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 16px;
  text-align: center;
}

.text_99 {
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 16px;
  text-align: center;
}

.line {
  margin-left: 46px;
  margin-right: 46px;
  opacity: 80%;
  width: 0px;
  height: 15.4px;
  border: 1px solid;
  border-color: #ffffff;
}

.xmxt {
  position: absolute;
  top: 182px;
}

.xmxt>div:not(:first-child) {
  margin-left: 85px;
}

.text_73 {
  margin-top: 15px;
  margin-bottom: 5px;
  font-family: Alibaba PuHuiTi;
  color: #181818;
  font-size: 28px;
}

.text_74 {
  margin-top: 5.5px;
  font-family: Alibaba PuHuiTi;
  color: #4d4d4d;
  font-size: 15px;
}

.sjx {
  width: 221px;
  height: 212px;
  background-image: linear-gradient(203.63deg,
      #c5f0fe 0%,
      rgba(210, 244, 241, 0.16) 19.24%,
      rgba(255, 255, 255, 0.43) 100%);
  border: 1px solid;
  border-color: #ffffff;
  border-radius: 6px;
}

.sjx:hover {
  width: 221px;
  height: 212px;
  background-image: linear-gradient(203.63deg,
      #c5f0fe 0%,
      #d2f4f1 19.24%,
      rgba(255, 255, 255, 0.43) 100%);
  border: 1px solid;
  border-color: #ffffff;
}

.sjx:hover .text_73 {
  margin-top: 15px;
  font-family: Alibaba PuHuiTi;
  color: #24a19b;
  font-size: 28px;
}

.sjx:hover .text_74 {
  margin-top: 5.5px;
  font-family: Alibaba PuHuiTi;
  color: #24a19b;
  font-size: 15px;
}

.pos_87 {
  position: absolute;
  bottom: 0px;
  background-color: rgba(255, 255, 255, 0.15);
  height: 108px;
  width: 100%;
  backdrop-filter: blur(20px);
}

.pos_88 {
  position: absolute;
  top: 3300px;
  height: 621px;
  width: 100%;
  background: url("/img/ros-hmci/底部.png");
  background-size: 100% 100%;
}

.pos_89 {
  margin-top: 20px;
  position: relative;
  top: 1180px;
  height: 100px;
}

.pos_80 {
  position: absolute;
  left: 50%;
  top: 80px;
  transform: translateX(-50%);
}

.pos_91 {
  position: absolute;
  top: 400px;
  right: 33px;
}

.pos_92 {
  position: absolute;
  top: 51px;
  right: 0px;
}

.pos_93 {
  position: absolute;
  top: -136px;
  right: -418px;
}


.pos_94 {
  position: absolute;
  top: -52px;
  right: -666px;
}

.pos_95 {
  position: absolute;
  top: 190px;
  right: -785px;
}

.pos_96 {
  position: absolute;
  top: 476px;
  right: -738px;
}

.pos_97 {
  position: absolute;
  left: -285px;
  top: 384px;
}

.pos_98 {
  position: absolute;
  left: -149px;
  top: 141px;
}

.pos_99 {
  position: absolute;
  left: 197px;
  top: 340px;
}

.section_37 {
  padding: 76px 0 52px;
  background-size: 100% 100%;
  background-repeat: no-repeat;
}

.space-y-50>*:not(:first-child) {
  margin-top: 50px;
}

.text_58 {
  color: #ffffff;
}

.text-wrapper_10 {
  padding: 18px 0;
  background-color: #e5b372;
  border-radius: 8px;
  width: 141px;
  border: solid 1px #00000000;
}

.text-wrapper_10:hover {
  background-color: #ebbd83;
  box-shadow: 0px 3px 12px rgba(30, 143, 138, 0.49);
}

.text_59 {
  color: #ffffff;
}

.section_38 {
  background-color: #1b2440;
}

.pos_55 {
  position: absolute;
  left: 0;
  right: 0;
  top: 3701.5px;
}

.group_29 {
  padding: 65px 0 62px;
  width: 1143px;
}

.space-y-16>*:not(:first-child) {
  margin-top: 16px;
}

.space-x-195>*:not(:first-child) {
  margin-left: 195px;
}

.space-y-32>*:not(:first-child) {
  margin-top: 32px;
}

.image-wrapper {
  padding: 22px 0 14px;
  background-color: #ffffff;
  border-radius: 13px;
  width: 100px;
}

.image_39 {
  width: 91px;
  height: 64px;
}

.text_66 {
  line-height: 15px;
}

.space-x-148>*:not(:first-child) {
  margin-left: 148px;
}

.font_19 {
  font-size: 18px;
  font-family: 苹方;
  line-height: 16.5px;
  color: #ffffff;
}

.font_20 {
  font-size: 14px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #bdc2d1;
}

.text_60 {
  margin-top: 28px;
}

.text_62 {
  margin-top: 22px;
}

.text_64 {
  margin-top: 22px;
}

.text_61 {
  margin-top: 28px;
}

.text_63 {
  margin-top: 22px;
}

.text_65 {
  margin-top: 22px;
}

.group_30 {
  margin-top: 2px;
}

.space-x-76>*:not(:first-child) {
  margin-left: 76px;
}

.space-y-26>*:not(:first-child) {
  margin-top: 26px;
}

.image_40 {
  border-radius: 4px;
  width: 99px;
  height: 100px;
}

.group_31 {
  height: 24px;
}

.text_68 {
  margin-top: 10px;
  color: #bdc2d1;
  font-size: 16px;
  font-family: PingFang SC;
  line-height: 14.5px;
}

.group_32 {
  margin-top: -24px;
  padding: 0 12px;
}

.space-x-118-reverse>*:not(:last-child) {
  margin-right: 118px;
}

.text_67 {
  line-height: 12.5px;
}

.section_39 {
  padding: 27px 0;
  background-color: #222a42;
}

.space-y-6>*:not(:first-child) {
  margin-top: 6px;
}

.group_33 {
  width: 395px;
  height: 12.5px;
}

.font_21 {
  font-size: 12px;
  font-family: PingFang SC;
  line-height: 13px;
  color: #bdc2d1;
}

.text_69 {
  line-height: 12.5px;
}

.text_70 {
  margin-top: -12px;
  line-height: 12.5px;
}

.text_71 {
  line-height: 14px;
  text-align: center;
  width: 381px;
}
</style>
