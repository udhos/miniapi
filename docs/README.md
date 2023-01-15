## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add miniapi https://udhos.github.io/miniapi

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo miniapi`
to see the charts.

To install the miniapi chart:

    helm install my-miniapi miniapi/miniapi
    #            ^          ^       ^
    #            |          |        \__ chart
    #            |           \__________ repo
    #             \_____________________ release (chart instance installed in cluster)

To uninstall the chart:

    helm delete my-miniapi

### Source

https://github.com/udhos/miniapi/tree/main/docs
