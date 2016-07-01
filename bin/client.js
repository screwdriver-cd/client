#!/usr/bin/env node
'use strict';

const vorpal = require('vorpal')();
const commander = require('commander');
const winston = require('winston');

const pkg = require('../package.json');
const pipelines = require('../lib/pipelines');
const jobs = require('../lib/jobs');
const builds = require('../lib/builds');
const config = require('../lib/config');

vorpal
    .command('pipelines create <scmUrl> [configUrl]')
    .description('Creates a new pipeline')
    .action((args) => pipelines.create(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('pipelines list')
    .description('List all pipelines')
    .action((args) => pipelines.list(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('jobs list')
    .description('List all jobs')
    .option('-p, --pipelineId <id>', 'Only show jobs that are for this pipeline')
    .action((args) => jobs.list(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('builds create <jobId>')
    .description('Creates a new build')
    .action((args) => builds.create(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('builds list')
    .description('List all builds')
    .option('-j, --jobId <id>', 'Only show builds that are for this job')
    .option('-s, --status <status>', 'Only show builds that have this status')
    .action((args) => builds.list(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('builds get <id>')
    .description('Get a specific build')
    .action((args) => builds.get(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('builds log <id>')
    .description('Get logs for a specific build')
    .action((args) => builds.log(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('config <field> [value]')
    .description('Get/Set local configuration')
    .action((args) => config.getSet(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

vorpal
    .command('login')
    .description('Get the URL to login to Screwdriver')
    .action((args) => config.login(args, vorpal.activeCommand.log.bind(vorpal.activeCommand)));

// Commander be down here
commander.version(pkg.version);
commander
    .command('pipelines-create')
    .description('Create a new pipeline')
    .arguments('<scmUrl> [configUrl]')
    .action((scmUrl, configUrl) => pipelines.create({ scmUrl, configUrl }, winston.info));

commander
    .command('pipelines-list')
    .description('List all pipelines')
    .action(() => pipelines.list({}, winston.info));

commander
    .command('jobs-list')
    .description('List all jobs')
    .option('-p, --pipelineId <id>', 'Only show jobs that are for this pipeline')
    .action((opts) => {
        const options = {
            pipelineId: opts.pipelineId
        };

        return jobs.list({ options }, winston.info);
    });

commander
    .command('builds-create')
    .description('Create a new build')
    .arguments('<jobId>')
    .action((jobId) => builds.create({ jobId }, winston.info));

commander
    .command('builds-list')
    .description('List all builds')
    .option('-j, --jobId <id>', 'Only show builds that are for this job')
    .option('-s, --status <status>', 'Only show builds that have this status')
    .action((opts) => {
        const options = {
            jobId: opts.jobId,
            status: opts.status
        };

        return builds.list({ options }, winston.info);
    });

commander
    .command('builds-get')
    .description('Get a specific build')
    .arguments('<id>')
    .action((id) => builds.get({ id }, winston.info));

commander
    .command('builds-log')
    .description('Get logs for a specific build')
    .arguments('<id>')
    .action((id) => builds.log({ id }, winston.info));

commander
    .command('config')
    .description('Get/Set local configuration')
    .arguments('<field> [value]')
    .action((field, value) => config.getSet({ field, value }, winston.info));

commander
    .command('login')
    .description('Get the URL to login to Screwdriver')
    .action(() => config.login({}, winston.info));

commander
    .command('repl')
    .description('switch into a CLI interface')
    .action(() => vorpal.delimiter('sd$').show());

commander.parse(process.argv);

if (!process.argv.slice(2).length) {
    commander.help();
}
