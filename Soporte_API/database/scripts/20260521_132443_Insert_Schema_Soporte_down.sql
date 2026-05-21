DELETE FROM soporte.reporte_mensajes
    WHERE mensaje IN (
        'El problema ocurre desde ayer, lo probé en iOS y Android también.',
        'Adjunto captura del extracto bancario con el doble cargo.',
        'Ya hablé con mi banco y me confirmaron que salieron dos débitos.'
    );

DELETE FROM soporte.reportes
    WHERE descripcion IN (
        'El botón de postulación no responde en móvil Chrome.',
        'Se me realizó un cobro duplicado por la suscripción Pro.',
        'Un usuario me envió mensajes ofensivos en el chat.'
    );