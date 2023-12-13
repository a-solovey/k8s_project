import dev.k8s.dir.gradle.KCommonLib
import dev.k8s.dir.gradle.KVersions
import dev.k8s.dir.gradle.SpringCloud

plugins {
    idea
    id("dev.k8s.dir.application.spring") version "1.1.1"
}

group = "dev.k8s.project"

k8s {
    val jooqKtxVersion = "1.1.1"

    using {
        KCommonLib.JOOQ_KTX.implementation(jooqKtxVersion)
        SpringCloud.CONFIG_SERVER.implementation(KVersions.springCloudVersion)

        "org.springframework.boot:spring-boot-starter-jdbc".implementation()
        "org.flywaydb:flyway-core".implementation()
        "io.kubernetes:client-java".implementation("18.0.1")
    }
}

val dbUrl: String by extra { System.getenv("DB_URL") ?: "jdbc:postgresql://localhost:5432/project" }
val dbUser: String by extra { System.getenv("DB_USER") ?: "postgres" }
val dbPassword: String by extra { System.getenv("DB_PASSWORD") ?: "password" }
val configEnv: String by extra { System.getenv("CONFIG_ENV") ?: "dev" }

tasks {
    val configImportTask = this.register("configImport", dev.k8s.project.ConfigImportTask::class.java)
    configImportTask.configure {
        yamlDirectory.fileValue(projectDir.resolve("configs"))
        jdbcUrl.set(dbUrl)
        username.set(dbUser)
        password.set(dbPassword)
    }
    val configExportTask = this.register("configExport", dev.k8s.project.ConfigExportTask::class.java)
    configExportTask.configure {
        yamlDirectory.fileValue(projectDir.resolve("configs"))
        jdbcUrl.set(dbUrl)
        username.set(dbUser)
        password.set(dbPassword)
        importEnv.set(configEnv)
    }
    val configExportSqlTask = this.register("configExportSql", dev.k8s.project.ConfigExportSqlTask::class.java)
    configExportSqlTask.configure {
        yamlDirectory.fileValue(projectDir.resolve("configs"))
        outFileName.set(projectDir.resolve("config-export-inserts.sql").apply { createNewFile() }.absolutePath)
        importEnv.set(configEnv)
    }
}