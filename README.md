# k8s_project - Проектная работа "Инфраструктурная платформа на основе Kubernetes"
Проектная работа представляет собой небольшую часть рабочего проекта, в связи с чем допущено ограничение - исходный код непосредственно ПО не выкладывается (не критично на мой взгляд), показаны примеры сборки и развертывания, также система логирования и мониторинга (такой вариант насколько было понятно из единственного вебинара по проектной работе, согласован руководителем курса и допускается). Ссылки на веб-интерфейс в открытом доступе нет, т.к. всё развертывание в закрытом контуре (возможно было бы показать трансляцией экрана, но т.к. опять же все вебинары для сдачи работ отменены, то в данном проекте предоставлены скриншоты, подтверждающие работающее развернутое ПО и мониторинг - в папке **screen** данного репозитория).

## Развертывание кластера Kubernetes и подготовка инфраструктуры
Подразумевается установка в закрытом контуре с использованием доступного в нем docker_registry (Nexus).  
Образы для инфраструкруры K8S заранее загружены в данный Nexus.  
Развертывание происходит через ansible. Также на данном этапе происходет развертывание и некоторая преднастройка таких компонентов, как, например, Istio.  
Всё для развертывания инфраструктуры находится в папке **infra** данного репозитория.  

Установка выполняется последовательным запуском двух плэйбуков:  
- ansible-playbook -i inventory/k8s_project/dev/inventory.ini k8s_api_lb.yml -vvv
- ansible-playbook -i inventory/k8s_project/dev/inventory.ini k8s_cluster.yml -vvv  

Для сброса всего кластера можно выполнить плэйбук:  
- ansible-playbook -i inventory/k8s_project/dev/inventory.ini k8s_reset.yml -vvv

## Развертывание прикладного ПО в кластере  
Усеченный проект включает в себя 3 модуля: **config-server**, **project-back**, **project-front**. Находятся в соответсвующих папках данного репозитория, каждая из которых имитирует отдельный репозиторий внутри гитлаба.  
Сборка в .gitlab-ci.yml: на шаге build выполняется сборка по тегу и выгружается образ в артефакты и в локально доступный Nexus.  
Деплой выполняется с помощью helm-чартов, которые находятся в отдельном репозитории, в данном проекте сымитировано в репозитории https://github.com/a-solovey/k8s_helm, также для необходимости развертывания в закрытом контуре без связи с гитлабом имеется скрипт deploy.sh для запуска деплоя, в целом развертывание реализовано в двух вариантах:  
### 1 вариант - сборка и автоматический деплой приложений в кластер локальный Kubernetes  
На шаге deploy в .gitlab-ci.yml через агент, связывающий гитлаб с локальным кластером K8S, происходит развертывание модуля, используя helm. В рабочем проекте планируется замена данного способа на Flux, на данный момент не реализовано.  
### 2 вариант - сборка и деплой приложений в кластер Kubernetes в закрытом контуре  
В этом варианте нет доступа к гитлабу и локальному Nexus со стороны закрытого контура, в обратную сторону нет доступа к кластеру K8S, в закрытом контуре доступен только свой Nexus (в примере указан везде один адрес Nexus, в рабочем проекте они отличаются), нексусы связаны между собой через проксирование. Развертывание выполняется локально в закрытом контуре вручную через запуск скрипта deploy.sh с указанием необходимых аргументов либо списком модулей, если необходимо развертнуть набор изменений в нескольких микросервисах.  

## Развертывание системы мониторинга и логирования  
Используется стек VictoriaMetrics, Grafana, OpenSearch. Всё необходимое находится в папке **monitoring** данного репозитория.  

Для установки сборщика метрик необходимо выполнить следующее (источник - папка helm):  
- helm install cert-manager ./cert-manager -n cert-manager --create-namespace --set installCRDs=true -f ./cert-manager/values-dev.yaml
- helm install opentelemetry-operator ./opentelemetry-operator -n opentelemetry --createnamespace -f ./opentelemetry-operator/values-dev.yaml
- helm install monitoring-collector ./monitoring-collector -n opentelemetry -f ./monitoringcollector/values-dev.yaml
- kubectl annotate namespace test instrumentation.opentelemetry.io/inject-java='myinstrumentation'  

Для установки дашбордов Grafana необходимо выполнить следующее (источник - папка ansible):  
- ansible-playbook -i inventory/gidnp/test/inventory.ini monitoring_config.yml

## Скриншоты по проекту  
Находятся в папке **screen** данного репозитория, скриншоты из Grafana при этом в основном взяты из реального рабочего проекта с бОльшим количеством микросервисов для наглядности.
