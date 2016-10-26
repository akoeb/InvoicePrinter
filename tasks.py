"""
This uses [http://docs.pyinvoke.org/](invoke) as command runner
"""
from invoke import task

args = '--rm --user 1000:1000 -v "${PWD}":/opt/app -v ${HOME}/.gitconfig:/.gitconfig -w /opt/app -e GOPATH=/opt/app'
img  = 'golang:1.6'

@task
def build(ctx):
    ctx.run('docker run {args} {img} go build -v'.format(args=args, img=img))


@task
def run(ctx):
    ctx.run('docker run {args} {img} go run *.go'.format(args=args, img=img))

@task
def get(ctx, pkg = ''):
    ctx.run('docker run {args} {img} go get -u -v {pkg}'.format(args=args, img=img, pkg=pkg))

@task
def pdf(ctx):
    ctx.run('docker run {args} {img} go run pdfwriter.go'.format(args=args, img=img))
